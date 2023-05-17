package service

import (
	"context"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) handleOrderCreated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaOrderCreated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	exists, err := r.orderExists(event.Order.OrderId.Int64())
	if err != nil {
		return errors.Wrap(err, "failed to check if order exists")
	}
	if exists {
		return nil
	}

	if err = r.addOrder(ctx, event.Order, event.UseRelayer); err != nil {
		return errors.Wrap(err, "failed to index order")
	}

	return nil
}

func (r *indexer) handleOrderUpdated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaOrderUpdated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	exists, err := r.orderExists(event.OrderId.Int64())
	if err != nil {
		return errors.Wrap(err, "failed to check if order exists")
	}
	if exists {
		return nil
	}

	if err = r.updateOrder(ctx, event.OrderId, event.Status); err != nil {
		return errors.Wrap(err, "failed to index order")
	}

	return nil
}

func (r *indexer) handleMatchCreated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaMatchCreated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	exists, err := r.matchExists(event.Match.MatchId.Int64())
	if err != nil {
		return errors.Wrap(err, "failed to check if match exists")
	}
	if exists {
		return nil
	}

	if err = r.addMatch(ctx, event.Match, event.UseRelayer); err != nil {
		return errors.Wrap(err, "failed to add match order")
	}

	return nil
}

func (r *indexer) handleMatchUpdated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaMatchUpdated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	exists, err := r.matchExists(event.MatchId.Int64())
	if err != nil {
		return errors.Wrap(err, "failed to check if match exists")
	}
	if exists {
		return nil
	}

	if err = r.updateMatch(ctx, event.MatchId, event.Status); err != nil {
		return errors.Wrap(err, "failed to update match order")
	}

	return nil
}
