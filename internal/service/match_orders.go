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

func (r *indexer) handleCreatedMatches(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := r.swapica.FilterMatchCreated(opts)
	if err != nil {
		return errors.Wrap(err, "failed to filter MatchCreated events")
	}
	for it.Next() {
		if err = r.addMatch(ctx, it.Event.Match); err != nil {
			return errors.Wrap(err, "failed to add match order")
		}

		if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
			r.lastBlock = b
			r.lastBlockOutdated = true
		}
	}

	return errors.Wrap(it.Error(), "error occurred while iterating over MatchCreated events")
}

func (r *indexer) handleUpdatedMatches(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := r.swapica.FilterMatchUpdated(opts, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter MatchUpdated events")
	}

	for it.Next() {
		if err = r.updateMatch(ctx, it.Event.MatchId, it.Event.Status); err != nil {
			return errors.Wrap(err, "failed to update match order")
		}

		if b := it.Event.Raw.BlockNumber + 1; b > r.lastBlock {
			r.lastBlock = b
			r.lastBlockOutdated = true
		}
	}

	return errors.Wrap(it.Error(), "error occurred while iterating over MatchUpdated events")
}

func (r *indexer) addMatch(ctx context.Context, mo gobind.ISwapicaMatch) error {
	log := r.log.WithField("match_id", mo.MatchId.String())
	log.Debug("adding new match order")
	body := requests.NewAddMatch(mo, r.chainID)
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
