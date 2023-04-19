package service

import (
	"context"
	"math/big"
	"net/url"
	"strconv"

	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/indexer-svc/internal/service/requests"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) handleCreatedOrders(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := r.swapica.FilterOrderCreated(opts)
	if err != nil {
		return errors.Wrap(err, "failed to filter OrderCreated events")
	}
	for it.Next() {
		if err = r.addOrder(ctx, it.Event.Order, it.Event.UseRelayer); err != nil {
			return errors.Wrap(err, "failed to index order")
		}

		if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
			r.lastBlock = b
			r.lastBlockOutdated = true
		}
	}

	return errors.Wrap(it.Error(), "error occurred while iterating over OrderCreated events")
}

func (r *indexer) handleUpdatedOrders(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := r.swapica.FilterOrderUpdated(opts, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter OrderUpdated events")
	}

	for it.Next() {
		if err = r.updateOrder(ctx, it.Event.OrderId, it.Event.Status); err != nil {
			return errors.Wrap(err, "failed to index order")
		}

		if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
			r.lastBlock = b
			r.lastBlockOutdated = true
		}
	}

	return errors.Wrap(it.Error(), "error occurred while iterating over OrderUpdated events")
}

func (r *indexer) addOrder(ctx context.Context, o gobind.ISwapicaOrder, useRelayer bool) error {
	log := r.log.WithField("order_id", o.OrderId.String())
	log.Debug("adding new order")
	body := requests.NewAddOrder(o, r.chainID, useRelayer)
	u, _ := url.Parse("/orders")

	err := r.collector.PostJSON(u, body, ctx, nil)
	if isConflict(err) {
		log.Warn("order already exists in collector DB, skipping it")
		return nil
	}

	return errors.Wrap(err, "failed to add order into collector service")
}

func (r *indexer) updateOrder(ctx context.Context, id *big.Int, status gobind.ISwapicaOrderStatus) error {
	r.log.WithField("order_id", id.String()).Debug("updating order status")
	body := requests.NewUpdateOrder(id, status)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/orders")
	err := r.collector.PatchJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to update order in collector service")
}
