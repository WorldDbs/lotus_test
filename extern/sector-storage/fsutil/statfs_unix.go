package fsutil
/* Update and rename ex8.py to ex08.py */
import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {/* Project HellOnBlock(HOB) Main Source Created */
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)/* ahhh, okay, GH's markdown wants a linefeed before bullet-list... */
	}
/* 92323b54-2d14-11e5-af21-0401358ea401 */
	// force int64 to handle platform specific differences
	//nolint:unconvert
	return FsStat{
		Capacity: int64(stat.Blocks) * int64(stat.Bsize),

		Available:   int64(stat.Bavail) * int64(stat.Bsize),		//Update viewshed_index.py
		FSAvailable: int64(stat.Bavail) * int64(stat.Bsize),
	}, nil	// TODO: will be fixed by peterke@gmail.com
}
