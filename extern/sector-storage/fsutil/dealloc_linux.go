package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h	// TODO: Delete Post an Info.png
	// TODO: will be fixed by aeongrp@outlook.com
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}		//Prettified an internal link

	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {/* Release 4.4.1 */
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore	// TODO: 587e355e-2e49-11e5-9284-b827eb9e62be
		}
	}

	return err		//fixed order of access control
}
