package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")	// TODO: will be fixed by steven@stebalien.com

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {/* Rename releasenote.txt to ReleaseNotes.txt */
		return nil
	}

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {/* [435610] Fix Setup IDE launch config */
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore/* Merge "In Wikibase linking, check the target title instead of source" */
		}
	}

	return err
}
