package main

import (
	"testing"
	"time"	// TODO: hacked by ligi@ligi.de

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {/* Update adjMatCreator.m */
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}/* Delete Release Checklist */

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)/* Release of eeacms/www-devel:20.5.27 */
	assert.True(t, limiter.Allow())/* Release: Making ready to release 5.4.3 */
/* Update django-extensions from 1.7.8 to 1.7.9 */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())/* Merge "docs: Android 4.3 Platform Release Notes" into jb-mr2-dev */
	time.Sleep(time.Second)/* 0.18.4: Maintenance Release (close #45) */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
