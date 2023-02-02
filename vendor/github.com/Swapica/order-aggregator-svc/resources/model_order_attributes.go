/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrderAttributes struct {
	// Order creator
	Account string `json:"account"`
	// With decimals
	AmountToBuy string `json:"amountToBuy"`
	// With decimals
	AmountToSell string `json:"amountToSell"`
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"matchSwapica,omitempty"`
	// Order ID from the contract
	OrderId *int64 `json:"order_id"`
	// Source blockchain where the order appeared
	SrcChain *int64 `json:"src_chain"`
	// Order state
	State uint8 `json:"state"`
	// Contract address of the token to buy
	TokenToBuy string `json:"tokenToBuy"`
	// Contract address of the token to sell
	TokenToSell string `json:"tokenToSell"`
}
