package bucket

import (
	"time"

	"github.com/go-zoox/kv"
)

// BucketKV is a in-BucketKV bucket.
type BucketKV struct {
	Storage kv.KV
	//
	maxAge time.Duration
}

// Config sets the max age of the bucket.
func (mb *BucketKV) Config(maxAge time.Duration) {
	mb.maxAge = maxAge
}

// Inc increases the counter by 1.
func (mb *BucketKV) Inc(key string) error {
	if !mb.Storage.Has(key) {
		expiresAt := time.Now().UnixNano()/int64(time.Millisecond) + int64(mb.maxAge/time.Millisecond)
		return mb.Storage.Set(key, &Value{
			Count:     1,
			ExpiresAt: expiresAt,
		}, mb.maxAge)
	}

	var value Value
	if err := mb.Storage.Get(key, &value); err != nil {
		return err
	}

	value.Count++

	return mb.Storage.Set(key, &value)
}

// Get returns the bucket value.
func (mb *BucketKV) Get(key string) (*Value, error) {
	var value Value
	if err := mb.Storage.Get(key, &value); err != nil {
		return nil, err
	}

	return &value, nil
}

// Del deletes the bucket value.
func (mb *BucketKV) Del(key string) error {
	return mb.Storage.Delete(key)
}
