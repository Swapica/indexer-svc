package service

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/Swapica/order-indexer-svc/internal/config"
	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/Swapica/order-indexer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	orderStateAwaitingMatch        uint8 = 1
	orderStateAwaitingFinalization uint8 = 2
	ethBlockPeriod                       = 10 * time.Second
)

type indexer struct {
	log       *logan.Entry
	swapica   *gobind.Swapica
	collector *jsonapi.Connector

	chainName      string
	requestTimeout time.Duration
	lastBlock      uint64

	blockURL, ordersURL, matchesURL *url.URL
}

func newIndexer(c config.Config, lastBlock uint64) indexer {
	chain := c.Network().ChainName
	block, err := url.Parse(chain + "/block")
	if err != nil {
		panic(errors.Wrap(err, "failed to parse URL"))
	}
	orders, _ := url.Parse(chain + "/orders")
	matches, _ := url.Parse(chain + "/match_orders")

	return indexer{
		log:            c.Log(),
		swapica:        c.Network().Swapica,
		collector:      c.Collector(),
		chainName:      chain,
		requestTimeout: c.Network().RequestTimeout,
		lastBlock:      lastBlock,
		blockURL:       block,
		ordersURL:      orders,
		matchesURL:     matches,
	}
}

func (r indexer) run(ctx context.Context) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()
	opts := &bind.FilterOpts{Context: childCtx, Start: r.lastBlock}

	it, err := r.swapica.FilterOrderUpdated(opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter OrderUpdated events")
	}

	changedLastBlock := false
	defer func() {
		if changedLastBlock {
			log := r.log.WithField("last_block", r.lastBlock)
			if err = r.saveLastBlock(ctx); err != nil {
				log.WithError(err).Error("failed to save last block")
			}
			log.Info("successfully saved last block")
		}
	}()

	// Warn: this logic may get stuck in various cases, think it over
	r.log.Debug("filtering OrderUpdated events")
	for {
		if !it.Next() {
			// Filtering events from the latest event's block instead of the latest confirmed one may be slower
			// However, an error may occur before the last event is saved, so all those next events will be lost
			if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
				r.lastBlock = b
				changedLastBlock = true
			}
			if err = it.Error(); err != nil {
				return errors.Wrap(err, "error occurred while filtering OrderUpdated events")
			}
			break
		}
		if err = r.indexOrder(ctx, it.Event); err != nil {
			return errors.Wrap(err, "failed to index order")
		}
	}

	matchIt, err := r.swapica.FilterMatchUpdated(opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter MatchUpdated events")
	}

	r.log.Debug("filtering MatchUpdated events")
	for {
		if !matchIt.Next() {
			r.lastBlock = matchIt.Event.Raw.BlockNumber
			if mb := matchIt.Event.Raw.BlockNumber; mb > r.lastBlock {
				r.lastBlock = mb + 1
				changedLastBlock = true
			}
			if err = matchIt.Error(); err != nil {
				return errors.Wrap(err, "error occurred while filtering MatchUpdated events")
			}
			break
		}
		if err = r.indexMatch(ctx, matchIt.Event); err != nil {
			return errors.Wrap(err, "failed to index match order")
		}
	}

	return nil
}

func (r indexer) saveLastBlock(ctx context.Context) error {
	body := resources.BlockResponse{
		Data: resources.Block{
			Key: resources.Key{
				ID:   strconv.FormatUint(r.lastBlock, 10),
				Type: resources.BLOCK,
			},
		},
	}

	err := r.collector.PostJSON(r.blockURL, body, ctx, nil)
	return errors.Wrap(err, "failed to set last block in collector")
}
