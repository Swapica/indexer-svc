package requests

import (
	"strconv"

	"github.com/Swapica/order-aggregator-svc/resources"
)

func NewUpdateBlock(number uint64) resources.BlockResponse {
	return resources.BlockResponse{
		Data: resources.Block{
			Key: resources.Key{
				ID:   strconv.FormatUint(number, 10),
				Type: resources.BLOCK,
			},
		},
	}
}
