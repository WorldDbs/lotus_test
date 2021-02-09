package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)/* Merge "Changing VPNaaS bug contact name" */
/* Release of eeacms/plonesaas:5.2.1-17 */
var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {	// TODO: hacked by ligi@ligi.de
	if length == 0 {/* Release: Splat 9.0 */
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)	// TODO: Updated view for camera, Removed translation implimentation.
	if errno, ok := err.(syscall.Errno); ok {		//Changed the stereo calibration command to not execute slam and gps
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {	// Add files by upload, hope to fix bug
			log.Warnf("could not deallocate space, ignoring: %v", errno)/* Added Release_VS2005 */
			err = nil // log and ignore
		}
	}

	return err
}
