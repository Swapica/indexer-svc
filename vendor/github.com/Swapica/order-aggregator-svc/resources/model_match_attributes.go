/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type MatchAttributes struct {
	// With decimals
	AmountToSell string `json:"amount_to_sell"`
	// Match order creator
	Creator string `json:"creator"`
	// Match order ID from the contract
	MatchId int64 `json:"match_id"`
	// ID of the order _from the contract_ to match
	OriginOrderId int64 `json:"origin_order_id"`
	// Match order state
	State uint8 `json:"state"`
	// Contract address of the token to sell
	TokenToSell string `json:"token_to_sell"`
}
