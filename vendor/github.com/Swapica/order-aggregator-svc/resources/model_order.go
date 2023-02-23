/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type Order struct {
	Key
	Attributes    OrderAttributes    `json:"attributes"`
	Relationships OrderRelationships `json:"relationships"`
}
type OrderResponse struct {
	Data     Order    `json:"data"`
	Included Included `json:"included"`
}

type OrderListResponse struct {
	Data     []Order         `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *OrderListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *OrderListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustOrder - returns Order from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOrder(key Key) *Order {
	var order Order
	if c.tryFindEntry(key, &order) {
		return &order
	}
	return nil
}
