/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type Block struct {
	Key
	// Only for generation of Go API resources
	Attributes *struct{} `json:"attributes,omitempty"`
}
type BlockResponse struct {
	Data     Block    `json:"data"`
	Included Included `json:"included"`
}

type BlockListResponse struct {
	Data     []Block         `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *BlockListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BlockListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBlock - returns Block from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBlock(key Key) *Block {
	var block Block
	if c.tryFindEntry(key, &block) {
		return &block
	}
	return nil
}
