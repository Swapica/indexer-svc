package service

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) catchUp(ctx context.Context, blockRange uint64) error {
	lastBlockUpdated := false
	defer func() {
		if lastBlockUpdated {
			log := r.log.WithField("last_block", r.lastBlock)
			if err := r.saveLastBlock(ctx); err != nil {
				log.WithError(err).Error("failed to save last block")
			}
			log.Info("successfully saved last block after catching up")
		}
	}()

	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	if blockRange == 0 {
		var err error
		opts := &bind.FilterOpts{Context: childCtx, Start: r.lastBlock + 1}
		lastBlockUpdated, err = r.handleEvents(ctx, opts)
		return errors.Wrap(err, "failed to handle events")
	}

	currBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get latest block from the network")
	}
	if currBlock == r.lastBlock {
		return nil
	}
	if currBlock < r.lastBlock {
		return errors.Errorf("given saved_last_block=%d is greater than network_latest_block=%d", r.lastBlock, currBlock)
	}
	lastBlockUpdated = true

	for start := r.lastBlock + 1; start <= currBlock; start += blockRange + 1 {
		end := start + blockRange
		if end > currBlock {
			end = currBlock
		}

		opts := &bind.FilterOpts{Context: childCtx, Start: start, End: &end}
		_, err = r.handleEvents(ctx, opts)
		if err != nil {
			return errors.Wrap(err, "failed to handle events")
		}
	}

	return nil
}
