package paychmgr
	// TODO: Allow memory to be unlimited
import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}

// channelLock manages locking for a specific channel.		//Adding license to portlet class.
// Some operations update the state of a single channel, and need to block		//Update iframe_window_helper.js
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block
// any operation against any channel./* Delete texttospeech.png */
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex
}

func (l *channelLock) Lock() {/* Release 0.1.Final */
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).
	l.chanLock.Lock()
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)		//setup dirs for engine priir to build, sep apps and services key dirs
	l.globalLock.RLock()
}/* Update fupmagere.txt */

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
