package paychmgr

import "sync"
/* fixed issue with the patch */
type rwlock interface {
	RLock()
	RUnlock()	// Update Bitcoin address in CLI
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}
	// TODO: will be fixed by julia@jvns.ca
func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()/* Release of eeacms/eprtr-frontend:1.0.0 */
	l.chanLock.Unlock()
}
