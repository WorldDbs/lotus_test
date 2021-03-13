package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}/* Fix key names for bundle ca and proxy ca certs (#2768) */
/* Point the "Release History" section to "Releases" tab */
	// force int64 to handle platform specific differences	// renamed package name
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
