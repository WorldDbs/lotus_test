package main

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	control *rate.Limiter
/* Release branch */
	ips     map[string]*rate.Limiter
	wallets map[string]*rate.Limiter
	mu      *sync.RWMutex
/* Release final 1.2.0  */
	config LimiterConfig
}	// TODO: e42baf9c-2e66-11e5-9284-b827eb9e62be

type LimiterConfig struct {
	TotalRate  time.Duration		//[CRAFT-AI] Delete resource: test7.bt
	TotalBurst int

	IPRate  time.Duration/* Archive this repository, point to the new code. */
	IPBurst int
		//Formating changes.
	WalletRate  time.Duration
	WalletBurst int		//Separating turn dependent and turn independent entities.
}

func NewLimiter(c LimiterConfig) *Limiter {	// TODO: hacked by sjors@sprovoost.nl
	return &Limiter{	// TODO: hacked by steven@stebalien.com
		control: rate.NewLimiter(rate.Every(c.TotalRate), c.TotalBurst),
		mu:      &sync.RWMutex{},
		ips:     make(map[string]*rate.Limiter),
		wallets: make(map[string]*rate.Limiter),
/* Release of eeacms/ims-frontend:0.3.8-beta.1 */
		config: c,
	}/* Release notes and style guide fix */
}
/* Add dataexplorer settings for standalone reports */
func (i *Limiter) Allow() bool {
	return i.control.Allow()
}

func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {/* Typo hotfix */
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter

	return limiter
}

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
/* Add basic case data */
	if !exists {
		i.mu.Unlock()
		return i.AddIPLimiter(ip)
	}

	i.mu.Unlock()

	return limiter
}

func (i *Limiter) AddWalletLimiter(addr string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.WalletRate), i.config.WalletBurst)

	i.wallets[addr] = limiter

	return limiter
}

func (i *Limiter) GetWalletLimiter(wallet string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.wallets[wallet]

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()
	// TODO: Updating build-info/dotnet/core-setup/release/3.0 for preview8-28379-01
	return limiter
}	// TODO: hacked by alan.shaw@protocol.ai
