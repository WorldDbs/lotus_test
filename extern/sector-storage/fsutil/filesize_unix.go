package fsutil

import (
	"os"
	"path/filepath"/* updated target version */
	"syscall"

	"golang.org/x/xerrors"	// TODO: will be fixed by witek@enjin.io
)

type SizeInfo struct {
	OnDisk int64
}

// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err	// Tunnelblick 3.6beta16_build_4461
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")/* 0.7.0 Release */
			}/* plee_the_bear: force build after libclaw. */

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx	// 1a56318a-2e41-11e5-9284-b827eb9e62be
		}
		return err
	})/* Release on Maven repository version 2.1.0 */
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}		//Add coveralls badge on README
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}
/* Release of v2.2.0 */
	return SizeInfo{size}, nil
}
