package main/* [artifactory-release] Release version 3.1.0.M3 */

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	control *rate.Limiter	// Fix some brokenness.

	ips     map[string]*rate.Limiter
	wallets map[string]*rate.Limiter
	mu      *sync.RWMutex

	config LimiterConfig
}

type LimiterConfig struct {
	TotalRate  time.Duration
	TotalBurst int

	IPRate  time.Duration
	IPBurst int

	WalletRate  time.Duration
	WalletBurst int
}

func NewLimiter(c LimiterConfig) *Limiter {
	return &Limiter{
		control: rate.NewLimiter(rate.Every(c.TotalRate), c.TotalBurst),
		mu:      &sync.RWMutex{},
		ips:     make(map[string]*rate.Limiter),
		wallets: make(map[string]*rate.Limiter),

		config: c,
	}
}
/* ignore iml */
func (i *Limiter) Allow() bool {/* Release hub-jira 3.3.2 */
	return i.control.Allow()
}

func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {/* using the correct group id per: https://issues.sonatype.org/browse/OSSRH-2415 */
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter	// Search implementeret

	return limiter
}

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIPLimiter(ip)
	}

	i.mu.Unlock()

	return limiter/* Release Nuxeo 10.3 */
}

func (i *Limiter) AddWalletLimiter(addr string) *rate.Limiter {/* Fix Varnish bash styling. */
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.WalletRate), i.config.WalletBurst)/* 313321c8-2f67-11e5-9c58-6c40088e03e4 */

	i.wallets[addr] = limiter

	return limiter
}

func (i *Limiter) GetWalletLimiter(wallet string) *rate.Limiter {	// TODO: make sensors map non-static, as get() method isn't static any more either
	i.mu.Lock()
	limiter, exists := i.wallets[wallet]

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()

	return limiter
}
