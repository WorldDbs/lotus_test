package fsutil

import (
	"os"
	"syscall"
	// TODO: Added new utils function in qFormat class (camelize, tableize...)
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")		//Merge "Hygiene: remove duplicate code in ListCardView"

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil/* Merge "Release 3.2.3.300 prima WLAN Driver" */
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}
	}

	return err
}
