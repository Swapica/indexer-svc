package service

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) catchUp(ctx context.Context, blockRange uint64) error {
	defer func() { r.updateLastBlock(ctx) }()

	if blockRange == 0 {
		var err error
		opts := &bind.FilterOpts{Start: r.lastBlock + 1}
		err = r.handleEvents(ctx, opts)
		return errors.Wrap(err, "failed to handle events")
	}

	currBlock, err := r.getNetworkLatestBlock(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get the latest block from the network")
	}
	if currBlock == r.lastBlock {
		return nil
	}

	// +1, because for eth_getLogs events with blockNumber == fromBlock or toBlock are included, so intersection is avoided
	for start := r.lastBlock + 1; start <= currBlock; start += blockRange + 1 {
		end := start + blockRange
		if end > currBlock {
			end = currBlock
		}

		opts := &bind.FilterOpts{Start: start, End: &end}
		if err = r.handleEvents(ctx, opts); err != nil {
			return errors.Wrap(err, "failed to handle events")
		}
	}

	r.lastBlock = currBlock
	r.lastBlockOutdated = true
	return nil
}
