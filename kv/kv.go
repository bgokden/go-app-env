package kv

type KV interface {
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Delete([]byte) error
}
