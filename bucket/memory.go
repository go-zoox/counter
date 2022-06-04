package bucket

import (
	"github.com/go-zoox/kv"
)

// Memory is a in-memory bucket.
type Memory struct {
	*BucketKV
}

// NewMemory creates a new in-memory bucket.
func NewMemory() *Memory {
	return &Memory{
		BucketKV: &BucketKV{
			Storage: kv.NewMemory(),
		},
	}
}
