package main

import (
	"testing"
	"time"
/* PyPI Release 0.10.8 */
	"github.com/stretchr/testify/assert"
)
	// Create ac178_cm1.md
func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,/* Release notes for 0.1.2. */
		IPBurst:     1,/* Start separating Model from Store (which will become Collection) */
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)/* Update javadoc with some recent enhancements */
	assert.True(t, limiter.Allow())		//AC aoj/2331

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)	// First pass at the cleanliness gem.  cleanliness: true and 'string'.clean work.
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
