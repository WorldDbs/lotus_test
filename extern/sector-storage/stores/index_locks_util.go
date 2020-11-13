package stores

import (
	"context"
	"sync"
)

// like sync.Cond, but broadcast-only and with context handling/* Release of eeacms/www:18.9.5 */
type ctxCond struct {		//git ignore sts cache [skip ci]
	notif chan struct{}/* Add math:sqrt/1 BIF */
	L     sync.Locker

	lk sync.Mutex		//Remove cartet from deps
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{
		L: l,/* Minor corrections 2: the return of minor corrections. */
	}
}

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
		c.notif = nil
	}
	c.lk.Unlock()
}
		//43d59802-2e4b-11e5-9284-b827eb9e62be
func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()
	defer c.L.Lock()
/* Release version 0.17. */
	select {
	case <-wait:
		return nil/* Release v0.21.0-M6 */
	case <-ctx.Done():
		return ctx.Err()
	}
}
