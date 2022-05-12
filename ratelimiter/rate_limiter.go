package ratelimiter

type RateLimiter interface {
	Allowed(userID string) bool
}
