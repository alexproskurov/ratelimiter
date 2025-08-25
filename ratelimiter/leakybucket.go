package ratelimiter

import (
	"sync"
	"time"
)

type leakyBucket struct {
	capacity  int
	rate      time.Duration
	available int
	lastLeak  time.Time
	mu        sync.Mutex
}

func NewLeakyBucket(capacity int, rate time.Duration) *leakyBucket {
	return &leakyBucket{
		capacity:  capacity,
		rate:      rate,
		available: 0,
		lastLeak:  time.Now(),
	}
}

func (b *leakyBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastLeak)
	leaked := int(elapsed / b.rate)

	if leaked > 0 {
		b.available -= leaked
		if b.available < 0 {
			b.available = 0
		}
		b.lastLeak = now
	}

	if b.available < b.capacity {
		b.available++
		return true
	}

	return false
}
