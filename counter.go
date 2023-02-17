package counter

import (
	"fmt"
	"time"

	"github.com/go-zoox/counter/bucket"
)

// Counter count total occurs during a period.
type Counter struct {
	prefix    string
	namespace string
	bucket    bucket.Bucket
}

// New creates a counter.
func New(bucket bucket.Bucket, namespace string, maxAge time.Duration) *Counter {
	bucket.Config(maxAge)

	return &Counter{
		bucket:    bucket,
		prefix:    "go-zoox:counter",
		namespace: namespace,
	}
}

func (c *Counter) key(id string) string {
	return fmt.Sprintf("%s:%s:%s", c.prefix, c.namespace, id)
}

// Inc increases the counter by 1.
func (c *Counter) Inc(id string) error {
	return c.bucket.Inc(c.key(id))
}

// Count counts the total number of occurrences during a period.
func (c *Counter) Count(id string) (*bucket.Value, error) {
	return c.bucket.Get(c.key(id))
}

// Reset resets the counter to 0.
func (c *Counter) Reset(id string) error {
	idx := c.key(id)
	if err := c.bucket.Del(idx); err != nil {
		return err
	}

	return c.bucket.Inc(idx)
}

// NewMemory creates a in-memory counter.
func NewMemory(namespace string, maxAge time.Duration) *Counter {
	return New(bucket.NewMemory(), namespace, maxAge)
}

// NewRedis creates a redis-based counter.
func NewRedis(namespace string, maxAge time.Duration, cfg *bucket.RedisConfig) (*Counter, error) {
	bucket, err := bucket.NewRedis(cfg)
	if err != nil {
		return nil, err
	}

	return New(bucket, namespace, maxAge), nil
}
