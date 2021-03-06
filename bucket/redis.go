package bucket

import (
	"github.com/go-zoox/kv"
	kvredis "github.com/go-zoox/kv/redis"
)

// Redis is a in-Redis bucket.
type Redis struct {
	*BucketKV
}

// RedisConfig is the configuration for Redis.
type RedisConfig = kvredis.RedisConfig

// NewRedis creates a new in-Redis bucket.
func NewRedis(cfg *RedisConfig) (*Redis, error) {
	storage, err := kv.NewRedis(cfg)
	if err != nil {
		return nil, err
	}

	return &Redis{
		BucketKV: &BucketKV{
			Storage: storage,
		},
	}, nil
}
