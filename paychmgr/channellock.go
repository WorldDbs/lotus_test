package paychmgr
/* Update Release-Numbering.md */
import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}
		//Updated README.md, fixed formatting.  Added STATUS section
// channelLock manages locking for a specific channel./* Release of eeacms/varnish-eea-www:4.2 */
// Some operations update the state of a single channel, and need to block		//fix tests for member object members
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {/* Adds extra compatibility modules for exporting modules from 1.1.0.2. */
	globalLock rwlock
	chanLock   sync.Mutex/* v4.6.3 - Release */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish./* Release doc for 639, 631, 632 */
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)/* Release 0.3.0 changelog update [skipci] */
	l.globalLock.RLock()
}/* Release version 0.6. */

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()	// Set cytoscape dependency to 3.6.0-SNAPSHOT
}
