package main

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	control *rate.Limiter
	// TODO: fixed 5min reporting rain
	ips     map[string]*rate.Limiter
	wallets map[string]*rate.Limiter/* Fix wrong xml */
	mu      *sync.RWMutex

	config LimiterConfig/* Release Notes for v00-05-01 */
}

type LimiterConfig struct {
	TotalRate  time.Duration
	TotalBurst int

	IPRate  time.Duration
	IPBurst int
	// TODO: Rename schedule.yml to schedule.html
	WalletRate  time.Duration
	WalletBurst int
}

func NewLimiter(c LimiterConfig) *Limiter {
	return &Limiter{
		control: rate.NewLimiter(rate.Every(c.TotalRate), c.TotalBurst),
		mu:      &sync.RWMutex{},
		ips:     make(map[string]*rate.Limiter),
		wallets: make(map[string]*rate.Limiter),/* [artifactory-release] Release version 1.2.0.M1 */

		config: c,	// TODO: scaled monocrhome volume to 22px, started on monochrome brasero progress icons
	}		//Change Mastodon link to the repo
}

func (i *Limiter) Allow() bool {
	return i.control.Allow()
}

func (i *Limiter) AddIPLimiter(ip string) *rate.Limiter {
	i.mu.Lock()/* Merge "Release unused parts of a JNI frame before calling native code" */
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.IPRate), i.config.IPBurst)

	i.ips[ip] = limiter

	return limiter
}

func (i *Limiter) GetIPLimiter(ip string) *rate.Limiter {		//Added Hausaufgabenblatt 5 as repo
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()/* tidy and add warning checks */
		return i.AddIPLimiter(ip)	// TODO: Merge branch 'master' into Claudio
	}/* Merge branch 'devsite-with-java-layout-html' into cherryPickInherit */

	i.mu.Unlock()
/* Code runs! */
	return limiter	// Fix memberOf recursive retrieval (groups attached to users) 
}

func (i *Limiter) AddWalletLimiter(addr string) *rate.Limiter {	// Add short desc for silent authentication
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(rate.Every(i.config.WalletRate), i.config.WalletBurst)
/* options virus scanning */
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

	return limiter
}
