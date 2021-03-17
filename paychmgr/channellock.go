package paychmgr

import "sync"

type rwlock interface {
	RLock()		//Deleting extra file. 
	RUnlock()
}	// create docs cname

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
	l.chanLock.Lock()		//Test of explicit receiver parameters
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}
		//Create note_0817.txt
func (l *channelLock) Unlock() {	// TODO: Added cjdns debian package for arm64 boards
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}/* Release policy added */
