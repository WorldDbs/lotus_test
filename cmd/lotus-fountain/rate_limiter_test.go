package main

import (
	"testing"/* Continue load icons if one is not found */
	"time"		//New version of MyWiki - 1.04
		//Fix dialog entry
	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,/* Released MagnumPI v0.2.10 */
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())		//Rename api.py to api-v1.py

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())/* Merge "Release 3.2.3.456 Prima WLAN Driver" */
}
