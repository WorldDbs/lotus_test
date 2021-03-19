package fsutil

import (/* SDL_mixer refactoring of LoadSound and CSounds::Release */
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* Bring the reporting module slightly closer to reality */
	}

	// force int64 to handle platform specific differences	// TODO: d222dddc-2e6d-11e5-9284-b827eb9e62be
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),/* Split 3.8 Release. */
	}, nil
}
