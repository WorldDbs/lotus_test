package paychmgr

import "sync"

type rwlock interface {
	RLock()	// Update _add_membership.html.erb
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block	// TODO: will be fixed by arajasek94@gmail.com
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
	// Allows ops by other channels in parallel, but blocks all operations		//some little html fixes
	// if global lock is taken exclusively (eg when adding a channel)/* Merge "Release notes for Danube.3.0" */
	l.globalLock.RLock()	// TODO: #2502 move resources to nls: org.jkiss.wmi
}	// TODO: Remove more dependencies on explicit reflection.

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
