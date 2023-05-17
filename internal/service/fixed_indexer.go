package service

import (
	"context"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Swapica/indexer-svc/internal/config"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/indexer-svc/internal/service/requests"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type indexerFixed struct {
	log       *logan.Entry
	swapica   *gobind.Swapica
	collector *jsonapi.Connector
	ethClient *ethclient.Client

	chainID           int64
	blockRange        uint64
	lastBlock         uint64
	lastBlockOutdated bool
	requestTimeout    time.Duration

	handlers   map[string]Handler
	swapicaAbi abi.ABI
}

func newIndexerFixed(c config.Config, lastBlock uint64) indexerFixed {
	swapicaAbi, err := abi.JSON(strings.NewReader(gobind.SwapicaMetaData.ABI))
	if err != nil {
		panic(errors.Wrap(err, "failed to get ABI"))
	}

	indexerInstance := indexerFixed{
		log:            c.Log(),
		swapica:        c.Network().Swapica,
		collector:      c.Collector(),
		ethClient:      c.Network().EthClient,
		chainID:        c.Network().ChainID,
		blockRange:     c.Network().BlockRange,
		lastBlock:      lastBlock,
		requestTimeout: c.Network().RequestTimeout,
		swapicaAbi:     swapicaAbi,
	}

	indexerInstance.handlers = map[string]Handler{
		"OrderCreated": indexerInstance.handleOrderCreated,
		"OrderUpdated": indexerInstance.handleOrderUpdated,
		"MatchCreated": indexerInstance.handleMatchCreated,
		"MatchUpdated": indexerInstance.handleMatchUpdated,
	}

	return indexerInstance
}

type Handler func(ctx context.Context, eventName string, log *types.Log) error

func (r *indexerFixed) handleOrderCreated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaOrderCreated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	if err = r.addOrder(ctx, event.Order, event.UseRelayer); err != nil {
		return errors.Wrap(err, "failed to index order")
	}

	// FIXME should we use this?
	if b := event.Raw.BlockNumber + 1; b > r.lastBlock {
		r.lastBlock = b
		r.lastBlockOutdated = true
	}

	return nil
}

func (r *indexerFixed) handleOrderUpdated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaOrderUpdated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	if err = r.updateOrder(ctx, event.OrderId, event.Status); err != nil {
		return errors.Wrap(err, "failed to index order")
	}

	// FIXME should we use this?
	if b := event.Raw.BlockNumber + 1; b > r.lastBlock {
		r.lastBlock = b
		r.lastBlockOutdated = true
	}

	return nil
}

func (r *indexerFixed) handleMatchCreated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaMatchCreated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	if err = r.addMatch(ctx, event.Match, event.UseRelayer); err != nil {
		return errors.Wrap(err, "failed to add match order")
	}

	// FIXME should we use this?
	if b := event.Raw.BlockNumber + 1; b > r.lastBlock {
		r.lastBlock = b
		r.lastBlockOutdated = true
	}

	return nil
}

func (r *indexerFixed) handleMatchUpdated(ctx context.Context, eventName string, log *types.Log) error {
	var event gobind.SwapicaMatchUpdated

	err := r.swapicaAbi.UnpackIntoInterface(&event, eventName, log.Data)
	if err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": eventName,
		})
	}

	if err = r.updateMatch(ctx, event.MatchId, event.Status); err != nil {
		return errors.Wrap(err, "failed to update match order")
	}

	// FIXME should we use this?
	if b := event.Raw.BlockNumber + 1; b > r.lastBlock {
		r.lastBlock = b
		r.lastBlockOutdated = true
	}

	return nil
}

func (r *indexerFixed) runFixed(ctx context.Context) error {
	// TODO handle unprocessed

	// TODO wait for events

	return nil
}

func (r *indexerFixed) waitForEvents(
	ctx context.Context, sub ethereum.Subscription, events <-chan types.Log,
) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-sub.Err():
			return errors.Wrap(err, "log subscription failed")
		case event := <-events:
			if err := r.handleEvent(ctx, event); err != nil {
				return errors.Wrap(err, "failed to handle event")
			}
		}
	}
}

func (r *indexerFixed) handleEvent(ctx context.Context, log types.Log) error {
	topic := log.Topics[0] // First topic must be a hashed signature of the event

	event, err := r.swapicaAbi.EventByID(topic)
	if err != nil {
		return errors.Wrap(err, "failed to get event by topic", logan.F{
			"topic": topic.Hex(),
		})
	}

	// TODO
	//processed, err := r.checkLogProcessed(&log)
	//if err != nil {
	//	return errors.Wrap(err, "failed to check if log is processed")
	//}
	//if processed { // just skip
	//	r.log.WithFields(logan.F{
	//		"event":   event.Name,
	//		"tx_hash": log.TxHash.Hex(),
	//	}).Debug("got already handled event")
	//	return nil
	//}

	handler, ok := r.handlers[event.Name]
	if !ok {
		return errors.From(errors.New("no handler for such event name"),
			logan.F{
				"event_name": event.Name,
			})
	}

	err = handler(ctx, event.Name, &log)
	return errors.Wrap(err, "handling of event failed", logan.F{
		"topic":      topic.Hex(),
		"event_name": event.Name,
	})
}

func (r *indexerFixed) addOrder(ctx context.Context, o gobind.ISwapicaOrder, useRelayer bool) error {
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

func (r *indexerFixed) updateOrder(ctx context.Context, id *big.Int, status gobind.ISwapicaOrderStatus) error {
	r.log.WithField("order_id", id.String()).Debug("updating order status")
	body := requests.NewUpdateOrder(id, status)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/orders")
	err := r.collector.PatchJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to update order in collector service")
}

func (r *indexerFixed) addMatch(ctx context.Context, mo gobind.ISwapicaMatch, useRelayer bool) error {
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

func (r *indexerFixed) updateMatch(ctx context.Context, id *big.Int, state uint8) error {
	r.log.WithField("match_id", id.String()).Debug("updating match state")
	body := requests.NewUpdateMatch(id, state)
	u, _ := url.Parse(strconv.FormatInt(r.chainID, 10) + "/match_orders")
	err := r.collector.PatchJSON(u, body, ctx, nil)
	return errors.Wrap(err, "failed to update match order in collector service")
}
