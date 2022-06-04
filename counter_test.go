package counter

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-zoox/counter/bucket"
)

func TestRateLimit(t *testing.T) {
	id := "127.0.0.1"
	r := New(bucket.NewMemory(), "web", 5*time.Second)

	if err := r.Inc(id); err != nil {
		t.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}

	if err := r.Inc(id); err != nil {
		t.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}

	if err := r.Inc(id); err != nil {
		t.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}

	if err := r.Reset(id); err != nil {
		t.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}

	time.Sleep(1 * time.Second)

	if err := r.Inc(id); err != nil {
		t.Fatal(err)
	}

	if v, err := r.Count(id); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("count:", v)
	}
}
