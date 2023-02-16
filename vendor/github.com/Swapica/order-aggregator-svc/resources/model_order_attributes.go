/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrderAttributes struct {
	// With decimals
	AmountToBuy string `json:"amount_to_buy"`
	// With decimals
	AmountToSell string `json:"amount_to_sell"`
	// Order creator
	Creator string `json:"creator"`
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"match_swapica,omitempty"`
	// Order ID from the contract
	OrderId *int64 `json:"order_id"`
	// Order state
	State uint8 `json:"state"`
	// Contract address of the token to buy
	TokenToBuy string `json:"token_to_buy"`
	// Contract address of the token to sell
	TokenToSell string `json:"token_to_sell"`
}
