package service

import (
	"context"
	"math/big"

	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/indexer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *indexer) indexMatch(ctx context.Context, evt *gobind.SwapicaMatchUpdated) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	status, err := r.swapica.MatchStatus(&bind.CallOpts{Context: childCtx}, evt.Id)
	if err != nil {
		return errors.Wrap(err, "failed to get match order status from contract")
	}
	log := r.log.WithFields(logan.F{"match_id": evt.Id, "match_state": status.State})

	if status.State == orderStateAwaitingFinalization {
		log.Debug("match order was created, trying to get and add it")
		err = r.addMatch(ctx, evt.Id, status)
		return errors.Wrap(err, "failed to add created match order")
	}
	if status.State == orderStateNone {
		log.Warn("found match order with state NONE, skipping it")
		return nil
	}

	log.Debug("match order must exist, updating its status")
	err = r.updateMatchStatus(ctx, evt.Id, status)
	return errors.Wrap(err, "failed to update match order status")
}

func (r *indexer) addMatch(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	childCtx, cancel := context.WithTimeout(ctx, r.requestTimeout)
	defer cancel()

	mo, err := r.swapica.Matches(&bind.CallOpts{Context: childCtx}, id)
	if err != nil {
		return errors.Wrap(err, "failed to get match order from contract")
	}

	body := resources.MatchResponse{
		Data: resources.Match{
			Key: resources.Key{
				ID:   mo.Id.String(),
				Type: resources.MATCH_ORDER,
			},
			Attributes: resources.MatchAttributes{
				Account:      mo.Account.String(),
				TokenToSell:  mo.TokenToSell.String(),
				AmountToSell: mo.AmountToSell,
				OriginChain:  mo.OriginChain,
				State:        status.State,
			},
			Relationships: resources.MatchRelationships{
				OriginOrder: resources.Relation{
					Data: &resources.Key{
						ID:   mo.OriginOrderId.String(),
						Type: resources.ORDER},
				},
			},
		},
	}

	err = r.collector.PostJSON(r.matchesURL, body, ctx, nil)
	if isConflict(err) {
		r.log.WithField("match_id", mo.Id.String()).
			Warn("match order already exists in collector DB, skipping it")
		return nil
	}
	return errors.Wrap(err, "failed to add match order into collector service")
}

func (r *indexer) updateMatchStatus(ctx context.Context, id *big.Int, status gobind.SwapicaStatus) error {
	body := resources.UpdateMatchRequest{
		Data: resources.UpdateMatch{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateMatchAttributes{
				State: status.State,
			},
		},
	}

	err := r.collector.PatchJSON(r.matchesURL, body, ctx, nil)
	return errors.Wrap(err, "failed to update match order in collector service")
}
