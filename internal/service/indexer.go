package service

import (
	"context"
	"net/url"
	"time"

	"github.com/Swapica/order-indexer-svc/internal/config"
	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/Swapica/order-indexer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const orderStatusAwaitingMatch uint8 = 1

type indexer struct {
	log       *logan.Entry
	swapica   *gobind.Swapica
	collector *jsonapi.Connector

	chainName      string
	requestTimeout time.Duration
	lastBlock      uint64

	blockURL, ordersURL *url.URL
}

func newIndexer(c config.Config, lastBlock uint64) indexer {
	chain := c.Network().ChainName
	block, err := url.Parse(chain + "/block")
	if err != nil {
		panic(errors.Wrap(err, "failed to parse URL"))
	}
	orders, _ := url.Parse(chain + "/orders")

	return indexer{
		log:            c.Log(),
		swapica:        c.Network().Swapica,
		collector:      c.Collector(),
		chainName:      c.Network().ChainName,
		requestTimeout: c.Network().RequestTimeout,
		lastBlock:      lastBlock,
		blockURL:       block,
		ordersURL:      orders,
	}
}

func (r indexer) run(ctx context.Context) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()
	opts := &bind.FilterOpts{Context: childCtx, Start: r.lastBlock}

	// TODO: also filter MatchUpdated
	// should I filter it all the time? Probably, I only need to create a subscription once, and it will do everything
	it, err := r.swapica.FilterOrderUpdated(opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter OrderUpdated event")
	}

	// Warn: this logic may get stuck in various cases, think it over
	for {
		if !it.Next() { // it should be done in the end, probably, I'm not sure
			r.lastBlock = it.Event.Raw.BlockNumber // will it be nil in this case?
			// not the best way, because I want to save the LATEST block if no events were found
			if err = it.Error(); err != nil {
				return errors.Wrap(err, "error occurred while filtering events")
			}
			// unsubscribe?
			break
		}
		if err = r.indexOrder(ctx, it.Event); err != nil {
			// var err error
			// err = ...
			// break
			return errors.Wrap(err, "failed to index order")
		}
	}
	// set last block wisely, when errors occur; try with defer
	return nil
}

func (r indexer) indexOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
	// If an order was created, index it, otherwise only update the status
	log := r.log.WithField("order_id", evt.Id)
	if evt.Status.State == orderStatusAwaitingMatch {
		log.Debug("order was created, trying to get and add it")
		err := r.addOrder(ctx, evt)
		return errors.Wrap(err, "failed to add created order")
	}

	log.Debug("order must exist, updating its status")
	err := r.updateOrderStatus(ctx, evt)
	return errors.Wrap(err, "failed to update order status")
}

func (r indexer) addOrder(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	o, err := r.swapica.Orders(&bind.CallOpts{Context: childCtx}, evt.Id)
	if err != nil {
		return errors.Wrap(err, "failed to get order from contract")
	}

	matchSwapica := evt.Status.MatchSwapica.String()
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
				ExecutedBy:   evt.Status.ExecutedBy,
				MatchSwapica: &matchSwapica,
				State:        evt.Status.State,
				TokenToBuy:   o.TokenToBuy.String(),
				TokenToSell:  o.TokenToSell.String(),
			},
		},
	}

	err = r.collector.PostJSON(r.ordersURL, body, ctx, nil)
	return errors.Wrap(err, "failed to add order into collector service")
}

func (r indexer) updateOrderStatus(ctx context.Context, evt *gobind.SwapicaOrderUpdated) error {
	matchSwapica := evt.Status.MatchSwapica.String()
	body := resources.UpdateOrderRequest{
		Data: resources.UpdateOrder{
			Key: resources.Key{
				ID:   evt.Id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateOrderAttributes{
				ExecutedBy:   evt.Status.ExecutedBy,
				MatchSwapica: &matchSwapica,
				State:        evt.Status.State,
			},
		},
	}

	err := r.collector.PatchJSON(r.ordersURL, body, ctx, nil)
	return errors.Wrap(err, "failed to update order in collector service")
}

// I would like to check that status is 204, but it's impossible when err==nil, and base connector handles errors well
//func (r indexer) sendRequest(ctx context.Context,path *url.URL, body, dst interface{}) error{
//}
