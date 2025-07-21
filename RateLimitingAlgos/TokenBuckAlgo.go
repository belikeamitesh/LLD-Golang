package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   float64
	fillRate   float64 // tokens per second
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(capacity, fillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		fillRate:   fillRate,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest(requestTokens float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens < requestTokens {
		return false
	}
	tb.tokens -= requestTokens
	return true
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.fillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}
