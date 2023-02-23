/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UpdateMatch struct {
	Key
	Attributes UpdateMatchAttributes `json:"attributes"`
}
type UpdateMatchRequest struct {
	Data     UpdateMatch `json:"data"`
	Included Included    `json:"included"`
}

type UpdateMatchListRequest struct {
	Data     []UpdateMatch   `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateMatchListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateMatchListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateMatch - returns UpdateMatch from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateMatch(key Key) *UpdateMatch {
	var updateMatch UpdateMatch
	if c.tryFindEntry(key, &updateMatch) {
		return &updateMatch
	}
	return nil
}
