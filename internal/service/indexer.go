package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Swapica/indexer-svc/internal/config"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
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

	chainID        int64
	lastBlock      uint64
	requestTimeout time.Duration
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
	var lastBlockUpdated bool
	defer func() {
		if lastBlockUpdated {
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

	lastBlockUpdated, err := r.handleEvents(ctx, opts)
	return errors.Wrap(err, "failed to handle events")
}

func (r *indexer) handleEvents(ctx context.Context, opts *bind.FilterOpts) (bool, error) {
	var lbu1, lbu2, lbu3, lbu4 bool
	lbu1, err := r.handleCreatedOrders(ctx, opts)
	if err != nil {
		return lbu1, errors.Wrap(err, "failed to handle created orders")
	}

	lbu2, err = r.handleCreatedMatches(ctx, opts)
	if err != nil {
		return lbu1 || lbu2, errors.Wrap(err, "failed to handle created match orders")
	}

	lbu3, err = r.handleUpdatedOrders(ctx, opts)
	if err != nil {
		return lbu1 || lbu2 || lbu3, errors.Wrap(err, "failed to handle updated orders")
	}

	lbu4, err = r.handleUpdatedMatches(ctx, opts)
	return lbu1 || lbu2 || lbu3 || lbu4, errors.Wrap(err, "failed to handle updated match orders")
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

	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/block")
	err := r.collector.PostJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to set last block in collector")
}

func isConflict(err error) bool {
	c, ok := err.(cerrors.Error)
	return ok && c.Status() == http.StatusConflict
}
