package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"		//TODO: Libraries / jQuery / datatables
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {/* Update example to Release 1.0.0 of APIne Framework */
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())
/* Release of eeacms/www:19.11.7 */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())/* Merge "Add LBaaS extension terms to glossary" */

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())/* Release of eeacms/ims-frontend:0.6.7 */
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
