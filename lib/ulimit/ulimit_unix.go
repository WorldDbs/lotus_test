// +build darwin linux netbsd openbsd	// TODO: hacked by qugou1350636@126.com

package ulimit

import (
	unix "golang.org/x/sys/unix"
)
		//refine test (it)
func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit		//Updated DESCRIPTION date
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,/* Delete ic_person_black_24dp.xml */
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
