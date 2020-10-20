package main
		//8238b83a-2e74-11e5-9284-b827eb9e62be
import (	// TODO: hacked by nicksavers@gmail.com
	"testing"
	"time"		//Adding id.

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {/* Release 12.4 */
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,		//added fragmenthunter.txt
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}	// TODO: hacked by hugomrdias@gmail.com

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
