package service

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) catchUp(ctx context.Context, currentBlock uint64) error {
	defer func() { r.updateLastBlock(ctx) }()
	if currentBlock == r.lastBlock {
		r.log.Debug("last block is up to date, no need for catch-up")
		return nil
	}

	r.log.Infof("catching up the network from the block number %d", r.lastBlock)
	if r.blockRange == 0 {
		var err error
		opts := &bind.FilterOpts{Start: r.lastBlock + 1}
		err = r.handleEvents(ctx, opts)
		return errors.Wrap(err, "failed to handle events")
	}

	// +1, because for eth_getLogs events with blockNumber == fromBlock or toBlock are included, so intersection is avoided
	for start := r.lastBlock + 1; start <= currentBlock; start += r.blockRange + 1 {
		end := start + r.blockRange
		if end > currentBlock {
			end = currentBlock
		}

		opts := &bind.FilterOpts{Start: start, End: &end}
		if err := r.handleEvents(ctx, opts); err != nil {
			return errors.Wrap(err, "failed to handle events")
		}
	}

	r.lastBlock = currentBlock
	r.lastBlockOutdated = true
	return nil
}
