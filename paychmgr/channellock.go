package paychmgr	// branching unstable (veqryn)

import "sync"

type rwlock interface {
	RLock()
	RUnlock()	// Delete server-pysrc.html
}

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block	// TODO: will be fixed by earlephilhower@yahoo.com
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel./* Released springjdbcdao version 1.7.25 */
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}
		//Improve code comments
func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish./* Create include.conf */
	// Exclusive per-channel (no other ops by this channel allowed)./* Add create critic */
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
