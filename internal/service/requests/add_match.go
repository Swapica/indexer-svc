package requests

import (
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewAddMatch(mo gobind.ISwapicaMatch, chainID int64) resources.MatchResponse {
	orderId := mo.MatchId.Int64()
	return resources.MatchResponse{
		Data: resources.Match{
			Key: resources.Key{
				Type: resources.MATCH_ORDER,
			},
			Attributes: resources.MatchAttributes{
				Account:      mo.Creator.String(),
				AmountToSell: mo.AmountToSell.String(),
				MatchId:      &orderId,
				SrcChain:     &chainID,
				State:        mo.State,
				TokenToSell:  mo.TokenToSell.String(),
			},
			Relationships: resources.MatchRelationships{
				OriginChain: resources.Relation{
					Data: &resources.Key{
						ID:   mo.OriginChainId.String(),
						Type: resources.CHAIN,
					},
				},
				OriginOrder: resources.Relation{
					Data: &resources.Key{
						ID:   mo.OriginOrderId.String(),
						Type: resources.ORDER,
					},
				},
			},
		},
	}
}
