// +build freebsd/* Remove training whitespace. */

package ulimit

import (
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)

func init() {	// TODO: Delete Dark Knight Custom Theme Sample.pdf
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit/* Merge branch 'master' into addtocartserializer */
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)/* Release date updated in comments */
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")/* Formerly make.h.~56~ */
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}

func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")
	}
	rlimit := unix.Rlimit{
		Cur: int64(soft),
		Max: int64(max),
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
