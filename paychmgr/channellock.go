package paychmgr

import "sync"
		//correlation plot picking code
type rwlock interface {
	RLock()/* Released 0.4.1 */
	RUnlock()/* 4.1.6-beta-12 Release Changes */
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex/* Update ReleaseManual.md */
}

func (l *channelLock) Lock() {		//updated version number for submitting to atmosphere
	// Wait for other operations by this channel to finish.		//Beta 5.1.01 added to windows title
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.	// TODO: will be fixed by timnugent@gmail.com
	// Allows ops by other channels in parallel, but blocks all operations	// TODO: Use length of children returned from RenderTree.childiter
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}/* Missing requires */
