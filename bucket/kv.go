package bucket

import (
	"time"

	"github.com/go-zoox/kv"
)

// KV is a in-KV bucket.
type KV struct {
	Storage kv.KV
	//
	maxAge time.Duration
}

// Config sets the max age of the bucket.
func (mb *KV) Config(maxAge time.Duration) {
	mb.maxAge = maxAge
}

// Inc increases the counter by 1.
func (mb *KV) Inc(key string) error {
	if !mb.Storage.Has(key) {
		expiresAt := time.Now().UnixNano()/int64(time.Millisecond) + int64(mb.maxAge/time.Millisecond)
		err := mb.Storage.Set(key, &Value{
			Count:     1,
			ExpiresAt: expiresAt,
		}, mb.maxAge)
		return err
	}

	var value Value
	if err := mb.Storage.Get(key, &value); err != nil {
		return err
	}

	value.Count++

	return mb.Storage.Set(key, &value)
}

// Get returns the bucket value.
func (mb *KV) Get(key string) (*Value, error) {
	var value Value
	if err := mb.Storage.Get(key, &value); err != nil {
		return nil, err
	}

	return &value, nil
}

// Del deletes the bucket value.
func (mb *KV) Del(key string) error {
	return mb.Storage.Delete(key)
}
