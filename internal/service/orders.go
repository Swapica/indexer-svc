package service

import (
	"context"
	"math/big"
	"net/url"
	"strconv"

	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) indexOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
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
	if status.State == orderStateNone {
		log.Warn("found order with state NONE, skipping it")
		return nil
	}

	log.Debug("order must exist, updating its status")
	err = r.updateOrderStatus(ctx, evt.Id, status)
	return errors.Wrap(err, "failed to update order status")
}

func (r *indexer) addOrder(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	o, err := r.swapica.Orders(&bind.CallOpts{Context: childCtx}, id)
	if err != nil {
		return errors.Wrap(err, "failed to get order from contract")
	}

	orderId := o.Id.Int64()
	body := resources.OrderResponse{
		Data: resources.Order{
			Key: resources.Key{
				Type: resources.ORDER,
			},
			Attributes: resources.OrderAttributes{
				Account:      o.Account.String(),
				AmountToBuy:  o.AmountToBuy.String(),
				AmountToSell: o.AmountToSell.String(),
				OrderId:      &orderId,
				SrcChain:     &r.chainID,
				State:        status.State,
				TokenToBuy:   o.TokenToBuy.String(),
				TokenToSell:  o.TokenToSell.String(),
			},
			Relationships: resources.OrderRelationships{
				DestChain: resources.Relation{
					Data: &resources.Key{
						ID:   o.DestChain.String(),
						Type: resources.CHAIN,
					},
				},
				ExecutedBy: nil,
			},
		},
	}

	u, _ := url.Parse("/orders")
	err = r.collector.PostJSON(u, body, ctx, nil)
	if isConflict(err) {
		r.log.WithField("order_id", o.Id.String()).
			Warn("order already exists in collector DB, skipping it")
		return nil
	}
	return errors.Wrap(err, "failed to add order into collector service")
}

func (r *indexer) updateOrderStatus(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	var matchSwapica *string
	if str := status.MatchSwapica.String(); str != ethAddress0 {
		matchSwapica = &str
	}

	var rel *resources.UpdateOrderRelationships
	// fixme: ensure it is not filled with 0 when it's empty
	if ex := status.ExecutedBy; ex == nil {
		rel = &resources.UpdateOrderRelationships{
			ExecutedBy: &resources.Relation{
				Data: &resources.Key{
					ID:   ex.String(),
					Type: resources.MATCH_ORDER,
				},
			}}
	}

	body := resources.UpdateOrderRequest{
		Data: resources.UpdateOrder{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateOrderAttributes{
				MatchSwapica: matchSwapica,
				State:        status.State,
			},
			Relationships: rel,
		},
	}

	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/orders")
	err := r.collector.PatchJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to update order in collector service")
}
