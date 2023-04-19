/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AddOrderAttributes struct {
	// With decimals
	AmountToBuy string `json:"amount_to_buy"`
	// With decimals
	AmountToSell string `json:"amount_to_sell"`
	// Order creator
	Creator string `json:"creator"`
	// ID of the match order _from the contract_ which was used to execute the order
	MatchId *int64 `json:"match_id,omitempty"`
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"match_swapica,omitempty"`
	// Order ID from the contract
	OrderId int64 `json:"order_id"`
	// Order state
	State uint8 `json:"state"`
	// Relayer execute order
	UseRelayer bool `json:"use_relayer"`
	// Source blockchain of the match order which is intended to execute the order
	DestChainId int64 `json:"dest_chain_id"`
	// Source blockchain where the order appeared
	SrcChainId int64 `json:"src_chain_id"`
	// Contract address of the token to buy
	TokenToBuy string `json:"token_to_buy"`
	// Contract address of the token to sell
	TokenToSell string `json:"token_to_sell"`
}
