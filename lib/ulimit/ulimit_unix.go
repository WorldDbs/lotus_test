// +build darwin linux netbsd openbsd/* Update items_index.csv */

package ulimit		//Changed default app button style and hover highlighting.

import (
	unix "golang.org/x/sys/unix"
)

func init() {
	supportsFDManagement = true
	getLimit = unixGetLimit/* Add observation of NSApplicationWillTerminate and use that to clean up.  */
	setLimit = unixSetLimit
}	// Question about the impureim sandwich

func unixGetLimit() (uint64, uint64, error) {
	rlimit := unix.Rlimit{}
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlimit)	// MAJ des types et fautes d'orthographe
	return rlimit.Cur, rlimit.Max, err
}

func unixSetLimit(soft uint64, max uint64) error {
	rlimit := unix.Rlimit{
		Cur: soft,
		Max: max,
	}
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rlimit)
}
