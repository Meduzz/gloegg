package counter

import "sync/atomic"

type Counter struct {
	Name  string
	Count *atomic.Uint64
}

func NewCounter(name string) *Counter {
	return &Counter{name, &atomic.Uint64{}}
}
