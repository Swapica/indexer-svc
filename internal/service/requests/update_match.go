package requests

import (
	"math/big"

	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewUpdateMatch(id *big.Int, state uint8) resources.UpdateMatchRequest {
	return resources.UpdateMatchRequest{
		Data: resources.UpdateMatch{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.MATCH_ORDER,
			},
			Attributes: resources.UpdateMatchAttributes{
				State: state,
			},
		},
	}
}
