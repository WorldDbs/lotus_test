package fsutil
/* Added DistributedQueue.peek(). */
import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)
	// Added nuget restore to pre build steps
type SizeInfo struct {
	OnDisk int64
}
/* Release documentation updates. */
// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {/* jueves 24 17:11 */
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}
	// TODO: hacked by denner@gmail.com
			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}		//Add zsh-command-time

	return SizeInfo{size}, nil		//Upgrade bijection to 0.7.2
}
