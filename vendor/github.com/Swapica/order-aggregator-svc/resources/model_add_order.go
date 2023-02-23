/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type AddOrder struct {
	Key
	Attributes AddOrderAttributes `json:"attributes"`
}
type AddOrderRequest struct {
	Data     AddOrder `json:"data"`
	Included Included `json:"included"`
}

type AddOrderListRequest struct {
	Data     []AddOrder      `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AddOrderListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AddOrderListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAddOrder - returns AddOrder from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAddOrder(key Key) *AddOrder {
	var addOrder AddOrder
	if c.tryFindEntry(key, &addOrder) {
		return &addOrder
	}
	return nil
}
