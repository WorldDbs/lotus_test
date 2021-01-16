package main/* packages: M7 updates */

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"		//android build of pktriggercord-cli
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,/* fde6c51a-2e5d-11e5-9284-b827eb9e62be */
		TotalBurst:  20,
		IPRate:      time.Second,/* Merge "Update RemoteController info when enabling/disabling it" into klp-dev */
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())/* nearby-handler stub added */
/* 309da5e8-2e75-11e5-9284-b827eb9e62be */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
