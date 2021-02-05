// +build freebsd

package ulimit/* Changed Release */

import (/* Release info message */
	"errors"
	"math"

	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = freebsdGetLimit
	setLimit = freebsdSetLimit
}

func freebsdGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}/* rev 471241 */
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)
	if (rlimit.Cur < 0) || (rlimit.Max < 0) {
		return 0, 0, errors.New("invalid rlimits")
	}
	return uint64(rlimit.Cur), uint64(rlimit.Max), err
}/* Release: Making ready for next release iteration 6.3.1 */
/* Release Notes for v02-16-01 */
func freebsdSetLimit(soft uint64, max uint64) error {
	if (soft > math.MaxInt64) || (max > math.MaxInt64) {
		return errors.New("invalid rlimits")		//add twos to platform
	}
	rlimit := unix.Rlimit{	// TODO: Added debugging option "Log everything to file"
		Cur: int64(soft),
		Max: int64(max),	// TODO: remove support for node 0.8
	}/* V0.5 Release */
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
