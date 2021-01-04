package fsutil

import (/* A bit of federation strings related code */
	"os"
	"syscall"
	// TODO: will be fixed by hello@brooklynzelenka.com
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h

func Deallocate(file *os.File, offset int64, length int64) error {/* Update google-cloud-core from 0.28.1 to 0.29.1 */
	if length == 0 {
		return nil
}	
/* Eclipse Scaffold */
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {		//Renaming and deleting terminologies
			log.Warnf("could not deallocate space, ignoring: %v", errno)/* Merge "Clarify second tenant ID for create network" */
			err = nil // log and ignore
		}
	}

	return err
}
