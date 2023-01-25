package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Swapica/order-indexer-svc/internal/config"
	"github.com/Swapica/order-indexer-svc/resources"
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
	var last uint64
	ctx := context.Background()

	// fixme: is runner really needed? And should it be separate function?
	running.UntilSuccess(ctx, s.log, "block-getter", func(_ context.Context) (bool, error) {
		raw := s.cfg.Network().ChainName + "/block"
		path, err := url.Parse(raw)
		if err != nil {
			return false, errors.Wrap(err, "failed to parse endpoint", map[string]interface{}{"url": raw})
		}

		var resp resources.BlockResponse
		if err = s.cfg.Collector().Get(path, &resp); err != nil {
			if err, ok := err.(cerrors.Error); ok && err.Status() == http.StatusNotFound {
				return true, nil
			}
			return false, errors.Wrap(err, "failed to get last block from collector")
		}

		last, err = strconv.ParseUint(resp.Data.ID, 10, 64)
		return err == nil, errors.Wrap(err, "failed to parse received block number", map[string]interface{}{"data.id": resp.Data.ID})
	}, time.Second, time.Hour)

	s.log.Infof("starting listening events from the block %d", last)
	runner := newIndexer(s.cfg, last)
	running.WithBackOff(ctx, s.log, "indexer", runner.run, time.Second, time.Second, time.Minute)

	return nil
}

func newService(cfg config.Config) *service {
	return &service{
		log: cfg.Log(),
		cfg: cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func (s *service) getLastBlock() (*uint64, error) {
	raw := s.cfg.Network().ChainName + "/block"
	path, err := url.Parse(raw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse endpoint", map[string]interface{}{"url": raw})
	}

	var resp resources.BlockResponse
	if err = s.cfg.Collector().Get(path, &resp); err != nil {
		if err, ok := err.(cerrors.Error); ok && err.Status() == http.StatusNotFound {
			var n uint64 = 0
			return &n, nil
		}
		return nil, errors.Wrap(err, "failed to get last block from collector")
	}

	n, err := strconv.ParseUint(resp.Data.ID, 10, 64)
	return &n, errors.Wrap(err, "failed to parse received block number", map[string]interface{}{"data.id": resp.Data.ID})
}
