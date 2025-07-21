package ratelimiter

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity      int           // Max requests that can be in the bucket
	leakRate      float64       // Requests per second (leak speed)
	bucket        []time.Time   // Stores timestamps of requests
	lastLeakTime  time.Time     // Last time we leaked
	mutex         sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		leakRate:     leakRate,
		bucket:       make([]time.Time, 0),
		lastLeakTime: time.Now(),
	}
}

func (lb *LeakyBucket) AllowRequest() bool {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	lb.leak()

	if len(lb.bucket) < lb.capacity {
		lb.bucket = append(lb.bucket, time.Now())
		return true
	}
	return false
}

func (lb *LeakyBucket) leak() {
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime).Seconds()
	leaked := int(elapsed * lb.leakRate)

	if leaked <= 0 {
		return
	}

	if leaked >= len(lb.bucket) {
		lb.bucket = lb.bucket[:0]
	} else {
		lb.bucket = lb.bucket[leaked:]
	}

	lb.lastLeakTime = now
}
