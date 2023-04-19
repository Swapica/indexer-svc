package requests

import (
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewAddOrder(o gobind.ISwapicaOrder, chainID int64, useRelayer bool) resources.AddOrderRequest {
	return resources.AddOrderRequest{
		Data: resources.AddOrder{
			Key: resources.Key{
				Type: resources.ORDER,
			},
			Attributes: resources.AddOrderAttributes{
				AmountToBuy:  o.AmountToBuy.String(),
				AmountToSell: o.AmountToSell.String(),
				UseRelayer:   useRelayer,
				Creator:      o.Creator.String(),
				MatchSwapica: nil, // must be nil by the Swapica contract
				OrderId:      o.OrderId.Int64(),
				State:        o.Status.State,
				TokenToBuy:   o.TokenToBuy.String(),
				TokenToSell:  o.TokenToSell.String(),
				DestChainId:  o.DestinationChain.Int64(),
				SrcChainId:   chainID,
			},
		},
	}
}
