package main	// Added access to generic properties in the results and custom metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
	// correct licence to GPL3.0
func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,		//ENH: add gaus function
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())	// Add support for y_iterate

	time.Sleep(time.Second)/* Release v0.0.8 */
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())		//main: fix return functions
	time.Sleep(time.Second)		//deps(varnish): update varnish to 6.4
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())/* Update PluginManager0001Test.php */
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
