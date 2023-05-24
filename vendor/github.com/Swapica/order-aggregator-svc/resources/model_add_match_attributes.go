/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AddMatchAttributes struct {
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
	// Relayer execute match order
	UseRelayer bool `json:"use_relayer"`
	// Source blockchain of the order to match
	OriginChainId int64 `json:"origin_chain_id"`
	// Source blockchain where the match order appeared
	SrcChainId int64 `json:"src_chain_id"`
	// Contract address of the token to sell
	TokenToSell string `json:"token_to_sell"`
}
