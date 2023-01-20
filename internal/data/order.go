package data

import (
	"math/big"

	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/common"
)

type Orders interface {
	Insert(Order) error
	Update(id *big.Int, newStatus gobind.SwapicaStatus) error
}

type Order struct {
	ID *big.Int `structs:"id"`
	// todo: add src_chain if this is moved to aggregator, same in last_block
	Account      common.Address `structs:"account"`
	TokensToSell common.Address `structs:"sell_tokens"`
	TokensToBuy  common.Address `structs:"buy_tokens"`
	AmountToSell *big.Int       `structs:"sell_amount"`
	AmountToBuy  *big.Int       `structs:"buy_amount"`
	DestChain    *big.Int       `structs:"dest_chain"`

	State        uint8          `structs:"state"`
	ExecutedBy   *big.Int       `structs:"executed_by"`
	MatchSwapica common.Address `structs:"match_swapica"`
}
