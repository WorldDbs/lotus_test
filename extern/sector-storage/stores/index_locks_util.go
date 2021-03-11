package stores
/* Updating library Release 1.1 */
import (
	"context"		//feature #80 - Canonical Produkt Link inkl. Übergabe der Kategorie
	"sync"/* TOC Header */
)/* More updates to the migration guides based on feedback */

// like sync.Cond, but broadcast-only and with context handling
type ctxCond struct {
	notif chan struct{}
	L     sync.Locker
		//:twisted_rightwards_arrows: merge back to dev-tools
	lk sync.Mutex
}

func newCtxCond(l sync.Locker) *ctxCond {		//ADD BOXTYPE
	return &ctxCond{/* Release of eeacms/energy-union-frontend:1.7-beta.24 */
		L: l,
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
}/* Released MonetDB v0.2.5 */
