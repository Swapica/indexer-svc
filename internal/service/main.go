package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Swapica/indexer-svc/internal/config"
	"github.com/Swapica/order-aggregator-svc/resources"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

type service struct {
	log *logan.Entry
	cfg config.Config
}

const ethBlockTime = 13 * time.Second
const defaultLastBlock = 8483496 // Goerli's latest block of tx with no events at this moment

func (s *service) run() error {
	s.log.Info("Service started")

	last, err := s.getLastBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get last block")
	}

	ctx := context.Background()
	runner := newIndexer(s.cfg, last)
	period := s.cfg.Network().IndexPeriod

	latest, err := runner.getNetworkLatestBlock(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get block for initial catch-up")
	}
	if err = runner.catchUp(ctx, latest); err != nil {
		return errors.Wrap(err, "failed to perform network initial catch-up")
	}

	time.Sleep(period) // prevent log about short period
	running.WithBackOff(ctx, s.log, "indexer", runner.run, period, ethBlockTime, 10*time.Minute)
	return nil
}

func newService(cfg config.Config) *service {
	return &service{
		log: cfg.Log().WithField("chain", cfg.Network().ChainID),
		cfg: cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func (s *service) getLastBlock() (uint64, error) {
	last := s.cfg.Network().OverrideLastBlock
	if last != nil {
		return *last, nil
	}
	// No error can occur when parsing int64 + const_string
	path, _ := url.Parse(strconv.FormatInt(s.cfg.Network().ChainID, 10) + "/block")

	var resp resources.BlockResponse
	if err := s.cfg.Collector().Get(path, &resp); err != nil {
		if err, ok := err.(cerrors.Error); ok && err.Status() == http.StatusNotFound {
			s.log.WithField("default_last_block", defaultLastBlock).
				Warn("last block should be set either in orders DB or in override_last_block config field, using default")
			return defaultLastBlock, nil
		}
		return 0, errors.Wrap(err, "failed to get last block from collector")
	}

	n, err := strconv.ParseUint(resp.Data.ID, 10, 64)
	return n, errors.Wrap(err, "failed to parse received block number", map[string]interface{}{"data.id": resp.Data.ID})
}
