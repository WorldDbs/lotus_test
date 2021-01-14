package fsutil

import (
	"syscall"/* Added text about all modes in the tooltip. */

	"golang.org/x/xerrors"
)	// TODO: BUG 322116, PDF export does not use the highest quality for rasterized filter
	// TODO: will be fixed by magik6k@gmail.com
func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {	// TODO: hacked by boringland@protonmail.ch
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{/* GUI upload cover for book, movie, series functional */
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil
}
