package service

import (
	"context"

	"github.com/Swapica/order-indexer-svc/internal/data"
	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const orderStatusAwaitingMatch uint8 = 1

func (s *service) worker(ctx context.Context) error {
	last, err := s.block.Get()
	if err != nil {
		return errors.Wrap(err, "failed to get last block")
	}
	s.log.Infof("start listening OrderUpdated events from the block %d", *last)

	events := make(chan *gobind.SwapicaOrderUpdated)
	childCtx, cancel := context.WithTimeout(ctx, s.network.RequestTimeout)
	defer cancel()
	opts := &bind.WatchOpts{Context: childCtx, Start: last}

	// TODO: replace with the filterer
	sub, err := s.network.WatchOrderUpdated(opts, events, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to OrderUpdated event")
	}

	for {
		select {
		case evt := <-events:
			if err = s.indexOrder(ctx, evt); err != nil {
				return errors.Wrap(err, "failed to index order")
			}
			s.log.WithField("order_id", evt.Id).Info("successfully indexed order")
		case err = <-sub.Err():
			return errors.Wrap(err, "subscription error occurred")
		case <-ctx.Done():
			return errors.New("context was cancelled")
		}
	}
}

func (s *service) indexOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
	// If an order was created, index it, otherwise only update the status
	// Warn: this logic may get stuck in various cases, think it over
	log := s.log.WithField("order_id", evt.Id)
	if evt.Status.State == orderStatusAwaitingMatch {
		log.Debug("order was created, trying to get it")
		o, err := s.getOrder(ctx, evt)
		if err != nil {
			return errors.Wrap(err, "failed to get created order")
		}

		if err = s.orders.Insert(o); err != nil {
			return errors.Wrap(err, "failed to add created order")
		}
	} else {
		log.Debug("order status was updated, updating ")
		if err := s.orders.Update(evt.Id, evt.Status); err != nil {
			return errors.Wrap(err, "failed to update order status")
		}
	}

	err := s.block.Set(evt.Raw.BlockNumber)
	return errors.Wrap(err, "failed to set new last block")
}

func (s *service) getOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) (data.Order, error) {
	childCtx, cancel := context.WithTimeout(ctx, s.network.RequestTimeout)
	defer cancel()

	o, err := s.network.Orders(&bind.CallOpts{Context: childCtx}, evt.Id)
	if err != nil {
		return data.Order{}, errors.Wrap(err, "failed to get order from contract")
	}

	return data.Order{
		ID:           o.Id,
		Account:      o.Account,
		TokensToSell: o.TokenToSell,
		TokensToBuy:  o.TokenToBuy,
		AmountToSell: o.AmountToSell,
		AmountToBuy:  o.AmountToBuy,
		DestChain:    o.DestChain,
		State:        evt.Status.State,
		ExecutedBy:   evt.Status.ExecutedBy,
		MatchSwapica: evt.Status.MatchSwapica,
	}, nil
}
