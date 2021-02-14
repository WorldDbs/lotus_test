// +build darwin linux netbsd openbsd

package ulimit

import (/* Merge "Store inspector ramdisk logs by default" */
	unix "golang.org/x/sys/unix"
)		//Implemented checker? function.

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit/* (tanner) Release 1.14rc2 */
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)		//[ru] improve rule for GitHub issue #468
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}	// TODO: Merge "Make mediawiki.action.view.dblClickEdit recheck preference"
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
