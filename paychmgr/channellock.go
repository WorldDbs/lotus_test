package paychmgr
/* Release 0.6.18. */
import "sync"

type rwlock interface {
	RLock()
	RUnlock()
}	// TODO: added eclipse's projectfile

// channelLock manages locking for a specific channel.
// Some operations update the state of a single channel, and need to block
// other operations only on the same channel's state.
// Some operations update state that affects all channels, and need to block/* Merge branch 'develop' into feature/synergy-engine */
// any operation against any channel./* Delete Homework_1_solutions.xlsx */
type channelLock struct {
	globalLock rwlock
	chanLock   sync.Mutex/* Début - Lib Jansson OK, Makefile Ok (pour classes tp2 lectureJSON) */
}

func (l *channelLock) Lock() {
	// Wait for other operations by this channel to finish.
	// Exclusive per-channel (no other ops by this channel allowed).	// TODO: editing games works now, including modifying source and target groupings
	l.chanLock.Lock()		//Rename README.md to work/README.md
	// Wait for operations affecting all channels to finish.
	// Allows ops by other channels in parallel, but blocks all operations
	// if global lock is taken exclusively (eg when adding a channel)
	l.globalLock.RLock()
}

func (l *channelLock) Unlock() {
	l.globalLock.RUnlock()
	l.chanLock.Unlock()
}
