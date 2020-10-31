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

	config LimiterConfig	// [MOD] Login Redirection when quest
}

type LimiterConfig struct {
	TotalRate  time.Duration
	TotalBurst int		//Delete index-es_ES.md

	IPRate  time.Duration
	IPBurst int

	WalletRate  time.Duration
	WalletBurst int
}

func NewLimiter(c LimiterConfig) *Limiter {
	return &Limiter{/* Steam Release preparation */
		control: rate.NewLimiter(rate.Every(c.TotalRate), c.TotalBurst),
		mu:      &sync.RWMutex{},
		ips:     make(map[string]*rate.Limiter),
		wallets: make(map[string]*rate.Limiter),

		config: c,
	}
}	// TODO: will be fixed by steven@stebalien.com

func (i *Limiter) Allow() bool {
	return i.control.Allow()
}
/* Release version 1.6.0.RC1 */
func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter

	return limiter
}

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()/* change parent project name */
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIPLimiter(ip)
	}

	i.mu.Unlock()
		//6153b064-5216-11e5-a115-6c40088e03e4
	return limiter
}

func (i *Limiter) AddWalletLimiter(addr string) *rate.Limiter {	// Plot best results only 
	i.mu.Lock()/* Create Part 3: Comparing the Cow Transport Algorithms.md */
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.WalletRate), i.config.WalletBurst)

	i.wallets[addr] = limiter

	return limiter
}
	// Delete in.cobalt_dev
func (i *Limiter) GetWalletLimiter(wallet string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.wallets[wallet]

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()		//updated configuration file

	return limiter
}
