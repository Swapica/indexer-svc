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

func (s *service) run() error {
	s.log.Info("Service started")

	last, err := s.getLastBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get last block")
	}

	runner := newIndexer(s.cfg, last)

	if s.cfg.Network().WsClient != nil {
		running.WithBackOff(
			context.Background(), s.log, "indexer",
			runner.run,
			s.cfg.Network().IndexPeriod, s.cfg.Network().IndexPeriod, time.Minute)
	} else {
		running.WithBackOff(
			context.Background(), s.log, "indexer",
			runner.runWithoutWs,
			s.cfg.Network().IndexPeriod, s.cfg.Network().IndexPeriod, time.Minute)
	}

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
	// No error can occur when parsing int64 + const_string
	path, _ := url.Parse(strconv.FormatInt(s.cfg.Network().ChainID, 10) + "/block")

	var resp resources.BlockResponse
	if err := s.cfg.Collector().Get(path, &resp); err != nil {
		if err, ok := err.(cerrors.Error); ok && err.Status() == http.StatusNotFound {
			s.log.WithField("default_last_block", s.cfg.Network().OverrideLastBlock).
				Warn("last block should be set either in orders DB or in override_last_block config field, using default")
			return s.cfg.Network().OverrideLastBlock, nil
		}
		return 0, errors.Wrap(err, "failed to get last block from collector")
	}

	n, err := strconv.ParseUint(resp.Data.ID, 10, 64)
	return n, errors.Wrap(err, "failed to parse received block number", map[string]interface{}{"data.id": resp.Data.ID})
}
