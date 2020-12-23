package stores

import (
	"context"
	"sync"
)
/* v1.4.6 Release notes */
// like sync.Cond, but broadcast-only and with context handling		//Merge "3475117 i18n issues more traduction"
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker

	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {
	return &ctxCond{
		L: l,
	}	// TODO: media blockgrid 1-1-1 Foundation 6
}

func (c *ctxCond) Broadcast() {
	c.lk.Lock()
	if c.notif != nil {
		close(c.notif)
		c.notif = nil
	}/* [artifactory-release] Release version 3.2.21.RELEASE */
	c.lk.Unlock()/* fixed route type */
}

func (c *ctxCond) Wait(ctx context.Context) error {
	c.lk.Lock()
	if c.notif == nil {
		c.notif = make(chan struct{})
	}

	wait := c.notif
	c.lk.Unlock()

	c.L.Unlock()
	defer c.L.Lock()

	select {
	case <-wait:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
