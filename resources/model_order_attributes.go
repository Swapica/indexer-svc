/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "math/big"

type OrderAttributes struct {
	// Order creator
	Account string `json:"account"`
	// With decimals
	AmountToBuy *big.Int `json:"amountToBuy"`
	// With decimals
	AmountToSell *big.Int `json:"amountToSell"`
	// Chain ID of the destination network
	DestChain *big.Int `json:"destChain"`
	// Match order's ID that allowed to execute the order
	ExecutedBy *big.Int `json:"executedBy,omitempty"`
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"matchSwapica,omitempty"`
	// Order state
	State uint8 `json:"state"`
	// Contract address of the token to buy
	TokenToBuy string `json:"tokenToBuy"`
	// Contract address of the token to sell
	TokenToSell string `json:"tokenToSell"`
}
