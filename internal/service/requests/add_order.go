package requests

import (
	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewAddOrder(o gobind.ISwapicaOrder, chainID int64) resources.OrderResponse {
	orderId := o.OrderId.Int64()
	return resources.OrderResponse{
		Data: resources.Order{
			Key: resources.Key{
				Type: resources.ORDER,
			},
			Attributes: resources.OrderAttributes{
				Account:      o.Creator.String(),
				AmountToBuy:  o.AmountToBuy.String(),
				AmountToSell: o.AmountToSell.String(),
				OrderId:      &orderId,
				SrcChain:     &chainID,
				State:        o.Status.State,
				TokenToBuy:   o.TokenToBuy.String(),
				TokenToSell:  o.TokenToSell.String(),
			},
			Relationships: resources.OrderRelationships{
				DestChain: resources.Relation{
					Data: &resources.Key{
						ID:   o.DestinationChain.String(),
						Type: resources.CHAIN,
					},
				},
				ExecutedBy: nil, // must be empty
			},
		},
	}
}
