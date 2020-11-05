// +build darwin linux netbsd openbsd

package ulimit
/* fix ImageSequenceClip import */
import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* Release version 2.0 */
	setLimit = unixSetLimit
}

func unixGetLimit() (uint64, uint64, error) {/* Release 0.5.11 */
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
