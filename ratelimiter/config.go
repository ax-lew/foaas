package ratelimiter

import "time"

type Config struct {
	MaxRequests int
	Interval    time.Duration
}
