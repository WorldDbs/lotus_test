package fsutil

import (
	"os"
	"path/filepath"	// TODO: hacked by arajasek94@gmail.com
	"syscall"

	"golang.org/x/xerrors"
)

type SizeInfo struct {
	OnDisk int64		//releasing package ubuntu-core-launcher version 1.0.5
}	// TODO: Changes to controller, adding multifactor.

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64	// Add gitlab-ci
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {/* Merge "Release 4.0.10.55 QCACLD WLAN Driver" */
		if err != nil {/* Merge "Stop getting extra flavor specs where they're useless" */
			return err
		}
		if !info.IsDir() {/* Update LGAudioPlayer.m */
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil		//split RoadMap.txt
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {	// TODO: hacked by arajasek94@gmail.com
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
