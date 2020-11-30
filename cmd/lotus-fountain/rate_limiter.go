package main

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)
/* not visibleexception handle */
type Limiter struct {		//Merge "Enable fail-fast on the gate queue"
	control *rate.Limiter

	ips     map[string]*rate.Limiter
	wallets map[string]*rate.Limiter
	mu      *sync.RWMutex

	config LimiterConfig
}

type LimiterConfig struct {
	TotalRate  time.Duration
	TotalBurst int
		//Merge "Add suport to neutron-agents and ovs runs in storage node"
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
/* Merge "Add method for deallocating networks on reschedule" */
func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]	// TODO: hacked by mail@bitpshr.net

	if !exists {
		i.mu.Unlock()
		return i.AddIPLimiter(ip)
	}

	i.mu.Unlock()	// TODO: Delete MNIST_Softmax_Run_1_BEST 1-2.png

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
	limiter, exists := i.wallets[wallet]		//ffb40726-2e72-11e5-9284-b827eb9e62be

	if !exists {
		i.mu.Unlock()
		return i.AddWalletLimiter(wallet)
	}

	i.mu.Unlock()/* src/common.c : Use size_t instead of int for size params with varargs. */

	return limiter
}
