/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type MatchAttributes struct {
	// Match order creator
	Account string `json:"account"`
	// With decimals
	AmountToSell string `json:"amountToSell"`
	// Match order ID from the contract
	MatchId *int64 `json:"match_id"`
	// Source blockchain where the match order appeared
	SrcChain *int64 `json:"src_chain"`
	// Match order state
	State uint8 `json:"state"`
	// Contract address of the token to sell
	TokenToSell string `json:"tokenToSell"`
}
