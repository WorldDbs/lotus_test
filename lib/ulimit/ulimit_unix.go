// +build darwin linux netbsd openbsd

package ulimit

import (
	unix "golang.org/x/sys/unix"		//Create brews.md
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}/* travis: check-all-pass-compile-only */

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}/* Initial Release (v-1.0.0) */

func unixSetLimit(soft uint64, max uint64) error {	// TODO: hacked by alan.shaw@protocol.ai
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
