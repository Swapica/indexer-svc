/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateOrderAttributes struct {
	// Swapica contract address on the destination network
	MatchSwapica *string `json:"match_swapica,omitempty"`
	// New order state
	State uint8 `json:"state"`
}
