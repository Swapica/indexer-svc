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

	requestTimeout time.Duration
	lastBlock      uint64

	blockURL, ordersURL, matchesURL *url.URL
}

func newIndexer(c config.Config, lastBlock uint64) indexer {
	chain := c.Network().ChainID
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
		requestTimeout: c.Network().RequestTimeout,
		lastBlock:      lastBlock,
		blockURL:       block,
		ordersURL:      orders,
		matchesURL:     matches,
	}
}

func (r *indexer) run(ctx context.Context) error {
	lastBlockUpd := false
	defer func() {
		if lastBlockUpd {
			log := r.log.WithField("last_block", r.lastBlock)
			if err := r.saveLastBlock(ctx); err != nil {
				log.WithError(err).Error("failed to save last block")
			}
			log.Info("successfully saved last block")
		}
	}()

	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()
	opts := &bind.FilterOpts{Context: childCtx, Start: r.lastBlock}

	lastBlockUpd, err := r.handleOrderEvents(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "failed to handle order events")
	}
	matchUpd, err := r.handleMatchEvents(ctx, opts)
	lastBlockUpd = lastBlockUpd || matchUpd
	return errors.Wrap(err, "failed to handle match order events")
}

func (r *indexer) saveLastBlock(ctx context.Context) error {
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

func (r *indexer) handleOrderEvents(ctx context.Context, opts *bind.FilterOpts) (lastBlockUpdated bool, err error) {
	it, err := r.swapica.FilterOrderUpdated(opts, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "failed to filter OrderUpdated events")
	}

	// Warn: this logic may get stuck in various cases, think it over
	for it.Next() {
		if err = r.indexOrder(ctx, it.Event); err != nil {
			return lastBlockUpdated, errors.Wrap(err, "failed to index order")
		}

		if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
			r.lastBlock = b
			lastBlockUpdated = true
		}
	}

	return lastBlockUpdated, errors.Wrap(it.Error(), "error occurred while iterating over OrderUpdated events")
}

func (r *indexer) handleMatchEvents(ctx context.Context, opts *bind.FilterOpts) (lastBlockUpdated bool, err error) {
	it, err := r.swapica.FilterMatchUpdated(opts, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "failed to filter MatchUpdated events")
	}

	for it.Next() {
		if err = r.indexMatch(ctx, it.Event); err != nil {
			return lastBlockUpdated, errors.Wrap(err, "failed to index match order")
		}

		if mb := it.Event.Raw.BlockNumber + 1; mb > r.lastBlock {
			r.lastBlock = mb
			lastBlockUpdated = true
		}
	}

	return lastBlockUpdated, errors.Wrap(it.Error(), "error occurred while iterating over MatchUpdated events")
}
