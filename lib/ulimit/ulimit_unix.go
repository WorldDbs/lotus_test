// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}	// Fixed line breaks (silly me)
/* Fix an issue with automount */
func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)		//(vila) Support MH-E in EmacsMail, using mml.
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//Fixed Issue 291 forceXMLOverwrite does not seem to work
}
