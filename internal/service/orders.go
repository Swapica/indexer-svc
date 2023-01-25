package service

import (
	"context"
	"math/big"

	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/Swapica/order-indexer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r indexer) indexOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	status, err := r.swapica.OrderStatus(&bind.CallOpts{Context: childCtx}, evt.Id)
	if err != nil {
		return errors.Wrap(err, "failed to get order status from contract")
	}
	log := r.log.WithFields(logan.F{"order_id": evt.Id, "order_state": status.State})

	if status.State == orderStateAwaitingMatch {
		log.Debug("order was created, trying to get and add it")
		err = r.addOrder(ctx, evt.Id, status)
		return errors.Wrap(err, "failed to add created order")
	}

	log.Debug("order must exist, updating its status")
	err = r.updateOrderStatus(ctx, evt.Id, status)
	return errors.Wrap(err, "failed to update order status")
}

func (r indexer) addOrder(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	o, err := r.swapica.Orders(&bind.CallOpts{Context: childCtx}, id)
	if err != nil {
		return errors.Wrap(err, "failed to get order from contract")
	}

	matchSwapica := status.MatchSwapica.String()
	body := resources.OrderResponse{
		Data: resources.Order{
			Key: resources.Key{
				ID:   o.Id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.OrderAttributes{
				Account:      o.Account.String(),
				AmountToBuy:  o.AmountToBuy,
				AmountToSell: o.AmountToSell,
				DestChain:    o.DestChain,
				ExecutedBy:   status.ExecutedBy,
				MatchSwapica: &matchSwapica,
				State:        status.State,
				TokenToBuy:   o.TokenToBuy.String(),
				TokenToSell:  o.TokenToSell.String(),
			},
		},
	}

	err = r.collector.PostJSON(r.ordersURL, body, ctx, nil)
	return errors.Wrap(err, "failed to add order into collector service")
}

func (r indexer) updateOrderStatus(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	matchSwapica := status.MatchSwapica.String()
	body := resources.UpdateOrderRequest{
		Data: resources.UpdateOrder{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateOrderAttributes{
				ExecutedBy:   status.ExecutedBy,
				MatchSwapica: &matchSwapica,
				State:        status.State,
			},
		},
	}

	err := r.collector.PatchJSON(r.ordersURL, body, ctx, nil)
	return errors.Wrap(err, "failed to update order in collector service")
}
