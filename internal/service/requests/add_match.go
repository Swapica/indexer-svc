package requests

import (
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewAddMatch(mo gobind.ISwapicaMatch, chainID int64) resources.AddMatchRequest {
	matchId := mo.MatchId.Int64()
	originChain := mo.OriginChainId.Int64()
	originOrder := mo.OriginOrderId.Int64()

	return resources.AddMatchRequest{
		Data: resources.AddMatch{
			Key: resources.Key{
				Type: resources.MATCH_ORDER,
			},
			Attributes: resources.AddMatchAttributes{
				AmountToSell:  mo.AmountToSell.String(),
				Creator:       mo.Creator.String(),
				MatchId:       &matchId,
				State:         mo.State,
				TokenToSell:   mo.TokenToSell.String(),
				OriginChainId: &originChain,
				OriginOrderId: &originOrder,
				SrcChainId:    &chainID,
			},
		},
	}
}
