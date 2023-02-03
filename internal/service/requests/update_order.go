package requests

import (
	"math/big"

	"github.com/Swapica/indexer-svc/internal/gobind"
	"github.com/Swapica/order-aggregator-svc/resources"
)

const ethAddress0 = "0x0000000000000000000000000000000000000000"

func NewUpdateOrder(id *big.Int, status gobind.ISwapicaOrderStatus) resources.UpdateOrderRequest {
	var matchSwapica *string
	if str := status.MatchSwapica.String(); str != ethAddress0 {
		matchSwapica = &str
	}

	var rel *resources.UpdateOrderRelationships
	// warn: ensure it is not filled with 0 when it's empty
	if ex := status.MatchId; ex == nil {
		rel = &resources.UpdateOrderRelationships{
			ExecutedBy: &resources.Relation{
				Data: &resources.Key{
					ID:   ex.String(),
					Type: resources.MATCH_ORDER,
				},
			}}
	}

	return resources.UpdateOrderRequest{
		Data: resources.UpdateOrder{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateOrderAttributes{
				MatchSwapica: matchSwapica,
				State:        status.State,
			},
			Relationships: rel,
		},
	}
}
