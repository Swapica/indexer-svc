package data

type LastBlock interface {
	Set(uint64) error
	Get() (*uint64, error)
}
