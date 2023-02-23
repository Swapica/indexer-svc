package requests

import (
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewAddMatch(m gobind.ISwapicaMatch, chainID int64) resources.AddMatchRequest {
	return resources.AddMatchRequest{
		Data: resources.AddMatch{
			Key: resources.Key{
				Type: resources.MATCH_ORDER,
			},
			Attributes: resources.AddMatchAttributes{
				AmountToSell:  m.AmountToSell.String(),
				Creator:       m.Creator.String(),
				MatchId:       m.MatchId.Int64(),
				State:         m.State,
				TokenToSell:   m.TokenToSell.String(),
				OriginChainId: m.OriginChainId.Int64(),
				OriginOrderId: m.OriginOrderId.Int64(),
				SrcChainId:    chainID,
			},
		},
	}
}
