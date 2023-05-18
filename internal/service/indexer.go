package service

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/Swapica/indexer-svc/internal/config"
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type indexer struct {
	log       *logan.Entry
	swapica   *gobind.Swapica
	collector *jsonapi.Connector
	ethClient *ethclient.Client
	wsClient  *ethclient.Client

	chainID           int64
	blockRange        uint64
	lastBlock         uint64
	lastBlockOutdated bool
	requestTimeout    time.Duration

	handlers        map[string]Handler
	swapicaAbi      abi.ABI
	contractAddress common.Address
}

type Handler func(ctx context.Context, eventName string, log *types.Log) error

func newIndexer(c config.Config, lastBlock uint64) indexer {
	swapicaAbi, err := abi.JSON(strings.NewReader(gobind.SwapicaMetaData.ABI))
	if err != nil {
		panic(errors.Wrap(err, "failed to get ABI"))
	}

	indexerInstance := indexer{
		log:             c.Log(),
		swapica:         c.Network().Swapica,
		collector:       c.Collector(),
		ethClient:       c.Network().EthClient,
		wsClient:        c.Network().WsClient,
		chainID:         c.Network().ChainID,
		lastBlock:       lastBlock,
		requestTimeout:  c.Network().RequestTimeout,
		swapicaAbi:      swapicaAbi,
		contractAddress: c.Network().ContractAddress,
	}

	indexerInstance.handlers = map[string]Handler{
		"OrderCreated": indexerInstance.handleOrderCreated,
		"OrderUpdated": indexerInstance.handleOrderUpdated,
		"MatchCreated": indexerInstance.handleMatchCreated,
		"MatchUpdated": indexerInstance.handleMatchUpdated,
	}

	return indexerInstance
}

func (r *indexer) run(ctx context.Context) error {
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block number")
	}

	newEvents := make(chan types.Log, 1024)
	sub, err := r.wsClient.SubscribeFilterLogs(ctx, r.filters(), newEvents)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to logs")
	}
	defer sub.Unsubscribe()

	if err := r.handleUnprocessedEvents(ctx, currentBlock); err != nil {
		return errors.Wrap(err, "failed to handle unprocessed events")
	}

	if err := r.waitForEvents(ctx, sub, newEvents); err != nil {
		return errors.Wrap(err, "failed to wait for unprocessed events")
	}

	return nil
}

func (r *indexer) handleUnprocessedEvents(
	ctx context.Context, currentBlock uint64,
) error {
	filters := r.filters()

	filters.FromBlock = new(big.Int).SetUint64(r.lastBlock)
	filters.ToBlock = new(big.Int).SetUint64(currentBlock + 1)

	logs, err := r.ethClient.FilterLogs(ctx, filters)
	if err != nil {
		return errors.Wrap(err, "failed to get filter logs")
	}

	for _, log := range logs {
		if err := r.handleEvent(ctx, log); err != nil {
			return errors.Wrap(err, "failed to handle event")
		}
	}

	return nil
}

func (r *indexer) waitForEvents(
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

func (r *indexer) handleEvent(ctx context.Context, log types.Log) error {
	topic := log.Topics[0] // First topic must be a hashed signature of the event

	event, err := r.swapicaAbi.EventByID(topic)
	if err != nil {
		return errors.Wrap(err, "failed to get event by topic", logan.F{
			"topic": topic.Hex(),
		})
	}

	handler, ok := r.handlers[event.Name]
	if !ok {
		return errors.From(errors.New("no handler for such event name"),
			logan.F{
				"event_name": event.Name,
			})
	}

	if err := handler(ctx, event.Name, &log); err != nil {
		return errors.Wrap(err, "handling of event failed", logan.F{
			"topic":      topic.Hex(),
			"event_name": event.Name,
		})
	}

	if err := r.updateLastBlock(ctx, log.BlockNumber); err != nil {
		return errors.Wrap(err, "failed to update last block")
	}

	return nil
}
