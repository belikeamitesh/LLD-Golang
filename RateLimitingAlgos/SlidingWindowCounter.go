package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	windowSizeInSeconds   int64
	maxRequestsPerWindow  float64
	currentWindowStartSec int64
	previousWindowCount   int64
	currentWindowCount    int64
	mu                    sync.Mutex
}

func NewSlidingWindowCounter(windowSize int64, maxRequests float64) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		windowSizeInSeconds:   windowSize,
		maxRequestsPerWindow:  maxRequests,
		currentWindowStartSec: time.Now().Unix(),
		previousWindowCount:   0,
		currentWindowCount:    0,
	}
}

func (s *SlidingWindowCounter) AllowRequest() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Unix()
	timePassed := now - s.currentWindowStartSec

	if timePassed >= s.windowSizeInSeconds {
		s.previousWindowCount = s.currentWindowCount
		s.currentWindowCount = 0
		s.currentWindowStartSec = now
		timePassed = 0
	}

	// Calculate weighted count
	weight := float64(s.windowSizeInSeconds-timePassed) / float64(s.windowSizeInSeconds)
	weightedCount := float64(s.previousWindowCount)*weight + float64(s.currentWindowCount)

// weightedCount = previousWindowCount * (1 - t / windowSize) + currentWindowCount
// Where:
// t = seconds passed in current window
// windowSize = total window size
// 1 - t/windowSize = how much of previous window still matters
// This is linear interpolation: the older the request, the less it matters.

	if weightedCount < s.maxRequestsPerWindow {
		s.currentWindowCount++
		return true
	}
	return false
}
