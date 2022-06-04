package bucket

import "time"

// Value is the bucket value
type Value struct {
	Count int64
	// the expires time, unit: milliseconds
	ExpiresAt int64
}

// Bucket is the butket to store the rate limit.
type Bucket interface {
	Config(maxAge time.Duration)
	Inc(key string) error
	Get(key string) (*Value, error)
	Del(key string) error
}
