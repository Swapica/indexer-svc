package service

import (
	"context"
	"net/url"
	"time"

	"github.com/Swapica/order-indexer-svc/internal/config"
	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	orderStateAwaitingMatch        uint8 = 1
	orderStateAwaitingFinalization uint8 = 2
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
		chainName:      c.Network().ChainName,
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
		return errors.Wrap(err, "failed to filter OrderUpdated event")
	}

	// Warn: this logic may get stuck in various cases, think it over
	r.log.Debug("filtering OrderUpdated events")
	for {
		if !it.Next() {
			r.lastBlock = it.Event.Raw.BlockNumber + 1
			// not the best way, because I want to save the LATEST block if no events were found
			if err = it.Error(); err != nil {
				return errors.Wrap(err, "error occurred while filtering OrderUpdated events")
			}
			break
		}
		if err = r.indexOrder(ctx, it.Event); err != nil {
			// var err error
			// err = ...
			// break
			return errors.Wrap(err, "failed to index order")
			// same notes are related to the following below
		}
	}

	matchIt, err := r.swapica.FilterMatchUpdated(opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter OrderUpdated event")
	}

	r.log.Debug("filtering MatchUpdated events")
	for {
		if !matchIt.Next() {
			r.lastBlock = matchIt.Event.Raw.BlockNumber
			if mb := matchIt.Event.Raw.BlockNumber; mb > r.lastBlock {
				r.lastBlock = mb + 1
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

	// set last block wisely, when errors occur; try with defer
	return nil
}
