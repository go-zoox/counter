package bucket

import (
	"github.com/go-zoox/kv"
)

// Memory is a in-memory bucket.
type Memory struct {
	*KV
}

// NewMemory creates a new in-memory bucket.
func NewMemory() *Memory {
	return &Memory{
		KV: &KV{
			Storage: kv.NewMemory(),
		},
	}
}
