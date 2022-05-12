package ratelimiter

import (
	"github.com/ax-lew/foaas/clock/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDefaultRateLimiter_Allowed(t *testing.T) {
	user := "user_1"
	otherUser := "user_2"

	clock := &mocks.Clock{}
	baseTime := time.Now()
	clock.On("Now").Return(baseTime).Once() // 1st request user_1
	clock.On("Now").Return(baseTime.Add(5 * time.Millisecond)).Once()
	clock.On("Now").Return(baseTime.Add(10 * time.Millisecond)).Once()
	clock.On("Now").Return(baseTime.Add(15 * time.Millisecond)).Once()
	clock.On("Now").Return(baseTime.Add(20 * time.Millisecond)).Once()            // 5th request user_1
	clock.On("Now").Return(baseTime.Add(20 * time.Millisecond)).Once()            // 1st request user_2
	clock.On("Now").Return(baseTime.Add(5*time.Second + time.Millisecond)).Once() // 6th request user_1

	limiter := NewLocalRateLimiter(&Config{
		MaxRequests: 4,
		Interval:    5 * time.Second,
	}, clock)

	// Fist 4 requests are allowed
	for i := 0; i < 4; i++ {
		allowed := limiter.Allowed(user)
		require.True(t, allowed)
	}
	// 5th request is not allowed
	allowed := limiter.Allowed(user)
	require.False(t, allowed)

	// Other user is allowed
	allowed = limiter.Allowed(otherUser)
	require.True(t, allowed)

	// 6th request is allowed
	allowed = limiter.Allowed(user)
	require.True(t, allowed)
}
