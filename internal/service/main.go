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

	ctx := context.Background()
	runner := newIndexer(s.cfg, last)

	s.log.Infof("catching up the network from the block number %d", last)
	if err = runner.catchUp(ctx, s.cfg.Network().BlockRange); err != nil {
		return errors.Wrap(err, "failed to catch up the network")
	}

	s.log.Infof("listening events in normal mode from the block number %d", runner.lastBlock)
	running.WithBackOff(ctx, s.log, "indexer", runner.run,
		s.cfg.Network().IndexPeriod, 5*time.Second, 10*time.Minute)

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
			return 0, errors.New("last block must be set either in orders database or in override_last_block config field")
		}
		return 0, errors.Wrap(err, "failed to get last block from collector")
	}

	n, err := strconv.ParseUint(resp.Data.ID, 10, 64)
	return n, errors.Wrap(err, "failed to parse received block number", map[string]interface{}{"data.id": resp.Data.ID})
}
