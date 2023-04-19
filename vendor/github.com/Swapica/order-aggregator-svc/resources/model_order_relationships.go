/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrderRelationships struct {
	DestinationChain Relation  `json:"destination_chain"`
	Match            *Relation `json:"match,omitempty"`
	SrcChain         Relation  `json:"src_chain"`
	TokenToBuy       Relation  `json:"token_to_buy"`
	TokenToSell      Relation  `json:"token_to_sell"`
}
