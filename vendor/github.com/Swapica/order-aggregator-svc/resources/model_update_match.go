/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateMatch struct {
	Key
	Attributes UpdateMatchAttributes `json:"attributes"`
}
type UpdateMatchRequest struct {
	Data     UpdateMatch `json:"data"`
	Included Included    `json:"included"`
}

type UpdateMatchListRequest struct {
	Data     []UpdateMatch `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
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
