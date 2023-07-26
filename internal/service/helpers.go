package service

import (
	"context"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/indexer-svc/internal/service/requests"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
)

var NotFound = errors.New("not found")

func (r *indexer) filters() ethereum.FilterQuery {
	topics := make([]common.Hash, 0, len(r.handlers))
	for eventName := range r.handlers {
		event := r.swapicaAbi.Events[eventName]

		topics = append(topics, event.ID)
	}

	filterQuery := ethereum.FilterQuery{
		Addresses: []common.Address{
			r.contractAddress,
		},
		Topics: [][]common.Hash{
			topics,
		},
	}
	return filterQuery
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

func (r *indexer) orderExists(id int64) (bool, error) {
	u, err := url.Parse(strconv.FormatInt(r.chainID, 10) + "/orders/" + strconv.FormatInt(id, 10)) // FIXME
	if err != nil {
		return false, errors.Wrap(err, "failed to parse url")
	}

	var order Order

	err = r.collector.Get(u, &order)
	if err != nil && err.Error() != NotFound.Error() {
		return false, errors.Wrap(err, "failed to get order")
	}

	return id == order.OrderID, nil
}

func (r *indexer) addMatch(ctx context.Context, mo gobind.ISwapicaMatch, useRelayer bool) error {
	log := r.log.WithField("match_id", mo.MatchId.String())
	log.Debug("adding new match order")
	body := requests.NewAddMatch(mo, r.chainID, useRelayer)
	u, _ := url.Parse("/match_orders")

	err := r.collector.PostJSON(u, body, ctx, nil)
	if isConflict(err) {
		log.Warn("match order already exists in collector DB, skipping it")
		return nil
	}

	return errors.Wrap(err, "failed to add match order into collector service")
}

func (r *indexer) updateMatch(ctx context.Context, id *big.Int, state uint8) error {
	r.log.WithField("match_id", id.String()).Debug("updating match state")
	body := requests.NewUpdateMatch(id, state)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/match_orders")
	err := r.collector.PatchJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to update match order in collector service")
}

func (r *indexer) matchExists(id int64) (bool, error) {
	u, err := url.Parse(strconv.FormatInt(r.chainID, 10) + "/match_orders/" + strconv.FormatInt(id, 10))
	if err != nil {
		return false, errors.Wrap(err, "failed to parse url")
	}

	var match Match

	err = r.collector.Get(u, &match)
	if err != nil && err.Error() != NotFound.Error() {
		return false, errors.Wrap(err, "failed to get match")
	}

	return id == match.MatchID, nil
}

func (r *indexer) updateLastBlock(ctx context.Context, lastBlock uint64) error {
	body := requests.NewUpdateBlock(lastBlock)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/block")
	err := r.collector.PostJSON(u, body, ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to save last block")
	}
	return nil
}

func isConflict(err error) bool {
	c, ok := err.(cerrors.Error)
	return ok && c.Status() == http.StatusConflict
}
