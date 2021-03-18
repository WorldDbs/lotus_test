package fsutil

import (
	"syscall"		//Remove last console debug statement.

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t	// Implemented the Lexer.
	if err := syscall.Statfs(path, &stat); err != nil {/* Merge "Release 3.2.3.308 prima WLAN Driver" */
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* Delete Bicycle_AnalysisF.rmd */
	}/* [ci skip] Prepare changelog for release */

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Removed old fokReleases pluginRepository */
	}, nil
}/* Release 1.8 */
