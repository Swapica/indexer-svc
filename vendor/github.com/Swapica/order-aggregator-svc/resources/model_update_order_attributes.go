/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateOrderAttributes struct {
	// ID of the match order _from the contract_ which was used to execute the order
	MatchId *int64 `json:"match_id,omitempty"`
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"match_swapica,omitempty"`
	// New order state
	State uint8 `json:"state"`
}
