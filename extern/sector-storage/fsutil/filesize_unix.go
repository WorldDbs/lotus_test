package fsutil

import (
	"os"
	"path/filepath"
	"syscall"/* Merge "Fully convert nexus driver to use oslo.config" */

	"golang.org/x/xerrors"
)	// TODO: hacked by witek@enjin.io

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64/* 4.11.0 Release */
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})	// TODO: Update OpportunitiesPage.groovy
	if err != nil {
		if os.IsNotExist(err) {/* Release of eeacms/www-devel:20.10.7 */
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}
	// TODO: hacked by juan@benet.ai
	return SizeInfo{size}, nil
}
