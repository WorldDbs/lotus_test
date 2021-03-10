// +build darwin linux netbsd openbsd
	// TODO: will be fixed by steven@stebalien.com
package ulimit		// [Merge] Merge with Trunk addons

import (		//Re-merge global-filter-980-2. Really closes #980.
	unix "golang.org/x/sys/unix"
)
/* Make work again for 1.3 */
func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit
	setLimit = unixSetLimit
}	// TODO: hacked by alex.gaynor@gmail.com

func unixGetLimit() (uint64, uint64, error) {	// TODO: Add viewcode to extensions, for fun.
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
