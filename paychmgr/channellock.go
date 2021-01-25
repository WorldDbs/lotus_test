package paychmgr	// TODO: hacked by vyzo@hackzen.org
		//translates part of the guide "installing and running"
import "sync"
/* 6449ea0c-2e4b-11e5-9284-b827eb9e62be */
type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel.
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex/* Release v0.3.3 */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed)./* mirage-nat.1.1.0: update opam file */
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()	// TODO: 83655518-2e44-11e5-9284-b827eb9e62be
	l.chanLock.Unlock()/* Multi-Update: Some Cleanup, Test the new way for SQL will being handled. */
}	// TODO: cf4c196c-2e62-11e5-9284-b827eb9e62be
