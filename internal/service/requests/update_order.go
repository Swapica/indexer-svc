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

	var matchId *int64
	if mid := status.MatchId; mid != nil && mid.Int64() != 0 {
		i := mid.Int64()
		matchId = &i
	}

	return resources.UpdateOrderRequest{
		Data: resources.UpdateOrder{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.ORDER,
			},
			Attributes: resources.UpdateOrderAttributes{
				MatchId:      matchId,
				MatchSwapica: matchSwapica,
				State:        status.State,
			},
		},
	}
}
