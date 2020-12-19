package main

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	control *rate.Limiter

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
}	// TODO: hacked by ng8eke@163.com

func (i *Limiter) Allow() bool {
	return i.control.Allow()
}

func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()/* Merge "Fix NPE when going "back" from Activity Transition." */
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter

	return limiter
}/* Added drl gui. */

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]	// Metadata tab: Delete config option added

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
	// Rails app Template ver. 1.1
	return limiter
}

func (i *Limiter) GetWalletLimiter(wallet string) *rate.Limiter {	// Formatting into columns
	i.mu.Lock()
	limiter, exists := i.wallets[wallet]

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()

	return limiter
}
