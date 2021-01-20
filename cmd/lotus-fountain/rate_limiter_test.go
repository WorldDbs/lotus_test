package main

import (
	"testing"
	"time"
		//Remove a couple months
	"github.com/stretchr/testify/assert"
)
	// TODO: hacked by 13860583249@yeah.net
func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,/* Add mogul interface header to sources. */
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)		//interface can extend only interface
	assert.True(t, limiter.Allow())	// Changed require_once to base_facebook.php
/* Release 2.8v */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
