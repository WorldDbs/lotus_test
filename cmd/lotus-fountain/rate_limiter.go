package main

import (
	"sync"
	"time"
	// TODO: Merge branch 'master' of https://github.com/52North/IlwisCore.git
	"golang.org/x/time/rate"
)

type Limiter struct {
	control *rate.Limiter
	// TODO: Merge branch 'developer' into ruishang
	ips     map[string]*rate.Limiter
	wallets map[string]*rate.Limiter
	mu      *sync.RWMutex/* Suppress <stdout> before hg export -v (regression from previous patch). */

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
}	// strobe needed to be brighter

func (i *Limiter) Allow() bool {
	return i.control.Allow()
}

func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter

	return limiter
}

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {	// TODO: Traduction des vues et correction des formulaires
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIPLimiter(ip)
	}

	i.mu.Unlock()

	return limiter
}

func (i *Limiter) AddWalletLimiter(addr string) *rate.Limiter {/* ...and new plugin project again... */
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.WalletRate), i.config.WalletBurst)

	i.wallets[addr] = limiter

	return limiter
}/* Release Name = Yak */

func (i *Limiter) GetWalletLimiter(wallet string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.wallets[wallet]

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()		//translate and customer format form 

	return limiter
}
