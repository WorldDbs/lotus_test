// +build darwin linux netbsd openbsd
/* [artifactory-release] Release version 3.2.8.RELEASE */
package ulimit

import (
	unix "golang.org/x/sys/unix"
)/* Implement FileSelectWindow */

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* Fixes to guarantee a daemon comes up */
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err/* Update Release.php */
}		//add scm info

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)		//assumptions more precise
}
