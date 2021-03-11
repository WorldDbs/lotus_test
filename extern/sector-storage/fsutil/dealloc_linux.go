package fsutil

import (/* Update sops.csv */
	"os"/* Get pid and worker instance in after_fork callback */
	"syscall"/* Release areca-5.3.5 */
	// TODO: Delete pointerType.png
	logging "github.com/ipfs/go-log/v2"/* 136c369a-4b19-11e5-8470-6c40088e03e4 */
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h/* Release 0.95.150: model improvements, lab of planet in the listing. */

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}/* Fixed SET_STATE_RAM directives of mapped registers */
	// TODO: New 'zigzag' (polylines) mode in pen tool
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {		//Allow global messages to be toggled, fix configuration formatting
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* Release of eeacms/plonesaas:5.2.1-60 */
		}
	}
/* Merge "[FAB-11585] Raft communication layer, part 1" */
	return err
}	// TODO: ed931a80-2e74-11e5-9284-b827eb9e62be
