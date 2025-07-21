package rate_limiter

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	windowSizeInSeconds   int64
	maxRequestsPerWindow  int64
	currentWindowStartSec int64
	requestCount          int64
	mu                    sync.Mutex
}

func NewFixedWindowCounter(windowSizeInSeconds, maxRequestsPerWindow int64) *FixedWindowCounter {
	return &FixedWindowCounter{
		windowSizeInSeconds:   windowSizeInSeconds,
		maxRequestsPerWindow:  maxRequestsPerWindow,
		currentWindowStartSec: time.Now().Unix(),
		requestCount:          0,
	}
}

func (f *FixedWindowCounter) AllowRequest() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	now := time.Now().Unix()
  
	if now-f.currentWindowStartSec >= f.windowSizeInSeconds {
		f.currentWindowStartSec = now
		f.requestCount = 0
	}

	if f.requestCount < f.maxRequestsPerWindow {
		f.requestCount++
		return true
	}

	return false 
}
