/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UpdateOrder struct {
	Key
	Attributes UpdateOrderAttributes `json:"attributes"`
}
type UpdateOrderRequest struct {
	Data     UpdateOrder `json:"data"`
	Included Included    `json:"included"`
}

type UpdateOrderListRequest struct {
	Data     []UpdateOrder   `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateOrderListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateOrderListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateOrder - returns UpdateOrder from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateOrder(key Key) *UpdateOrder {
	var updateOrder UpdateOrder
	if c.tryFindEntry(key, &updateOrder) {
		return &updateOrder
	}
	return nil
}
