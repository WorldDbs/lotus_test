// +build darwin linux netbsd openbsd
	// TODO: hacked by lexy8russo@outlook.com
package ulimit

import (
	unix "golang.org/x/sys/unix"
)
/* SEMPERA-2846 Release PPWCode.Kit.Tasks.NTServiceHost 3.3.0 */
func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)		//More unit tests + helper classes
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
