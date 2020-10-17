package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"		//Compatibility for old DualIso sessions
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h/* Changed evaluation logic to Google Page Speed API (3) */

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {/* Added debug prints */
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)		//clean up and kill some warnings
			err = nil // log and ignore
		}
	}

	return err
}
