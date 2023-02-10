package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Swapica/indexer-svc/internal/config"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/indexer-svc/internal/service/requests"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type indexer struct {
	log       *logan.Entry
	swapica   *gobind.Swapica
	collector *jsonapi.Connector
	ethClient *ethclient.Client

	chainID           int64
	lastBlock         uint64
	lastBlockOutdated bool
	requestTimeout    time.Duration
}

func newIndexer(c config.Config, lastBlock uint64) indexer {
	return indexer{
		log:            c.Log(),
		swapica:        c.Network().Swapica,
		collector:      c.Collector(),
		ethClient:      c.Network().EthClient,
		chainID:        c.Network().ChainID,
		lastBlock:      lastBlock,
		requestTimeout: c.Network().RequestTimeout,
	}
}

func (r *indexer) run(ctx context.Context) error {
	defer func() { r.updateLastBlock(ctx) }()

	// ensure that on the next iteration it filters precisely from currBlock+1 to neither skip orders, nor get duplicates
	currBlock, err := r.getNetworkLatestBlock(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get the latest block from the network")
	}

	opts := &bind.FilterOpts{Start: r.lastBlock, End: &currBlock}
	if err = r.handleEvents(ctx, opts); err != nil {
		return errors.Wrap(err, "failed to handle events")
	}

	r.lastBlock = currBlock + 1 // see catch_up.go
	r.lastBlockOutdated = true
	return nil
}

func (r *indexer) handleEvents(ctx context.Context, opts *bind.FilterOpts) error {
	child, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()
	opts.Context = child

	if err := r.handleCreatedOrders(ctx, opts); err != nil {
		return errors.Wrap(err, "failed to handle created orders")
	}

	if err := r.handleCreatedMatches(ctx, opts); err != nil {
		return errors.Wrap(err, "failed to handle created match orders")
	}

	if err := r.handleUpdatedOrders(ctx, opts); err != nil {
		return errors.Wrap(err, "failed to handle updated orders")
	}

	err := r.handleUpdatedMatches(ctx, opts)
	return errors.Wrap(err, "failed to handle updated match orders")
}

func (r *indexer) getNetworkLatestBlock(ctx context.Context) (uint64, error) {
	child, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	n, err := r.ethClient.BlockNumber(child)
	if err != nil {
		return n, errors.Wrap(err, "failed to get eth_blockNumber")
	}
	if n < r.lastBlock {
		return n, errors.Errorf("given saved_last_block=%d is greater than network_latest_block=%d", r.lastBlock, n)
	}

	return n, nil
}

func (r *indexer) updateLastBlock(ctx context.Context) {
	log := r.log.WithField("last_block", r.lastBlock)
	if !r.lastBlockOutdated {
		log.Debug("no updates of the last block")
		return
	}

	body := requests.NewUpdateBlock(r.lastBlock)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/block")
	err := r.collector.PostJSON(u, body, ctx, nil)
	if err != nil {
		log.WithError(err).Error("failed to save last block")
		return
	}
	r.lastBlockOutdated = false
	log.Debug("successfully saved last block")
}

func isConflict(err error) bool {
	c, ok := err.(cerrors.Error)
	return ok && c.Status() == http.StatusConflict
}
