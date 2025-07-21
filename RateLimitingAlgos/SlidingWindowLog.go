package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLog struct {
	windowSize          int64      // in seconds
	maxRequests         int
	requestTimestamps   []int64
	mu                  sync.Mutex
}

func NewSlidingWindowLog(windowSize int64, maxRequests int) *SlidingWindowLog {
	return &SlidingWindowLog{
		windowSize:        windowSize,
		maxRequests:       maxRequests,
		requestTimestamps: []int64{},
	}
}

func (sw *SlidingWindowLog) AllowRequest() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now().Unix()
	windowStart := now - sw.windowSize

	// Remove old timestamps outside the window
	filtered := []int64{}
	for _, ts := range sw.requestTimestamps {
		if ts > windowStart {
			filtered = append(filtered, ts)
		}
	}
	sw.requestTimestamps = filtered

	if len(sw.requestTimestamps) < sw.maxRequests {
		sw.requestTimestamps = append(sw.requestTimestamps, now)
		return true
	}
	return false
}
