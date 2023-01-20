package service

import (
	"context"
	"time"

	"github.com/Swapica/order-indexer-svc/internal/config"
	"github.com/Swapica/order-indexer-svc/internal/data"
	"github.com/Swapica/order-indexer-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

type service struct {
	log     *logan.Entry
	network config.Network
	block   data.LastBlock
	orders  data.Orders
}

func (s *service) run() error {
	s.log.Info("Service started")
	running.WithBackOff(context.Background(), s.log, "main", s.worker, time.Second, time.Second, time.Minute)

	return nil
}

func newService(cfg config.Config) *service {
	chain := cfg.Network().ChainName
	block, err := postgres.NewLastBlock(cfg.DB(), chain)
	if err != nil {
		panic(errors.Wrap(err, "failed to instantiate last block DB API"))
	}

	return &service{
		log:     cfg.Log(),
		network: cfg.Network(),
		block:   block,
		orders:  postgres.NewOrders(cfg.DB(), chain),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
