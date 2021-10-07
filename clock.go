package clock

import (
	"runtime"
	"time"
)

type Clock interface {
	After(d time.Duration) <-chan time.Time
}

// Sleep momentarily so that other goroutines can process.
func gosched() { runtime.Gosched() }

func New() Clock {
	return &clock{}
}

type clock struct{}

func (c *clock) After(d time.Duration) <-chan time.Time { return time.After(d) }
