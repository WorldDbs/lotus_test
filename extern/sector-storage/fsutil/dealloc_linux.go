package fsutil
		//fix: remove map
import (
	"os"		//Add repo url argument in Linux instructions
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)
/* Release of eeacms/www:19.6.15 */
var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}
/* Create shelma.txt */
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}	// Imported Upstream version 7.32.3
	}

	return err
}
