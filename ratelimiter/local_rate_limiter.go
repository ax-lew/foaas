package ratelimiter

import (
	"github.com/ax-lew/foaas/clock"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type LocalRateLimiter struct {
	mutex sync.Mutex
	usersRequests *cache.Cache
	config        *Config
	clock         clock.Clock
}

func NewLocalRateLimiter(config *Config, clock clock.Clock) *LocalRateLimiter {
	// We can remove the entries from the cache that haven't been updated in the last interval
	usersRequests := cache.New(config.Interval, 2*config.Interval)
	return &LocalRateLimiter{
		config:        config,
		clock:         clock,
		mutex:         sync.Mutex{},
		usersRequests: usersRequests,
	}
}

func (l *LocalRateLimiter) Allowed(userID string) bool {
	l.mutex.Lock()
	r, ok := l.usersRequests.Get(userID)
	var requests []time.Time
	if ok {
		requests = r.([]time.Time)
	} else {
		requests = make([]time.Time, 0)
	}

	now := l.clock.Now()
	requests = l.clearExpired(now, requests)

	l.mutex.Unlock()

	if len(requests) >= l.config.MaxRequests {
		return false
	}
	requests = append(requests, now)
	l.usersRequests.Set(userID, requests, 0)
	return true
}

func (l *LocalRateLimiter) clearExpired(now time.Time, requests []time.Time) []time.Time {
	validRequests := make([]time.Time, 0)
	for _, r := range requests {
		if now.Sub(r) > l.config.Interval {
			continue
		}
		validRequests = append(validRequests, r)
	}
	return validRequests
}
