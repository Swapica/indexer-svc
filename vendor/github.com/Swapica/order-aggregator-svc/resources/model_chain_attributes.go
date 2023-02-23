/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChainAttributes struct {
	ChainParams EvmChainParams `json:"chain_params"`
	// Type of blockchain by supported wallets, APIs, etc.
	ChainType string `json:"chain_type"`
	// Link to network icon
	Icon         *string `json:"icon,omitempty"`
	Name         string  `json:"name"`
	SwapContract string  `json:"swap_contract"`
}
