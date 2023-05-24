/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type AddMatch struct {
	Key
	Attributes AddMatchAttributes `json:"attributes"`
}
type AddMatchRequest struct {
	Data     AddMatch `json:"data"`
	Included Included `json:"included"`
}

type AddMatchListRequest struct {
	Data     []AddMatch      `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AddMatchListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AddMatchListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAddMatch - returns AddMatch from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAddMatch(key Key) *AddMatch {
	var addMatch AddMatch
	if c.tryFindEntry(key, &addMatch) {
		return &addMatch
	}
	return nil
}
