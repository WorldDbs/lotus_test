package fsutil

import (/* Repositioned anchor elements so they’ll validate */
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}	// Generated site for typescript-generator-maven-plugin 2.14.506

	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),
		//agrego la ayuda
		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),	// TODO: Update Readme.MD to say gradle instead of sbt
	}, nil
}
