// +build darwin linux netbsd openbsd/* remove old functions / variables */

package ulimit
/* Update SpreadsheetViewTable */
import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true/* Flexibilizando o método each. */
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
rre ,xaM.timilr ,ruC.timilr nruter	
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
