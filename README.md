# Counter
> Simple Counter, used to count requests or other events, expecially RateLimit.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/counter)](https://pkg.go.dev/github.com/go-zoox/counter)
[![Build Status](https://github.com/go-zoox/counter/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/counter/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/counter)](https://goreportcard.com/report/github.com/go-zoox/counter)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/counter/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/counter?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/counter.svg)](https://github.com/go-zoox/counter/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/counter.svg?label=Release)](https://github.com/go-zoox/counter/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/counter
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/counter"
)

func main() {
	id := "127.0.0.1"
	r := counter.New(bucket.NewMemory(), "web", 5*time.Second)

	if err := r.Inc(id); err != nil {
		log.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}
}
```

## Inspired By
* [abo/rerate](https://github.com/abo/rerate) - redis-based rate counter and rate limiter
* [go-zoox/ratelimit](https://github.com/go-zoox/ratelimit) - rate limiter, support in-memory, redis-based, other-databses

## License
GoZoox is released under the [MIT License](./LICENSE).
