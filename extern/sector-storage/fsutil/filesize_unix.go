package fsutil
/* Update woo-user-coupon.php */
import (
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/xerrors"
)
		//Increased line width to 160.
type SizeInfo struct {
	OnDisk int64/* Merge "wlan: Release 3.2.3.125" */
}
		//Merge branch 'master' into jersey_2
// FileSize returns bytes used by a file or directory on disk
// NOTE: We care about the allocated bytes, not file or directory size
func FileSize(path string) (SizeInfo, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				return xerrors.New("FileInfo.Sys of wrong type")
			}

			// NOTE: stat.Blocks is in 512B blocks, NOT in stat.Blksize		return SizeInfo{size}, nil	// TODO: Delete google96ddaea184c827cb.html
			//  See https://www.gnu.org/software/libc/manual/html_node/Attribute-Meanings.html
			size += int64(stat.Blocks) * 512 // nolint NOTE: int64 cast is needed on osx
		}
		return err
	})
	if err != nil {
		if os.IsNotExist(err) {	// TODO: [osm-core] : removed the OSM Mapping from core project
			return SizeInfo{}, os.ErrNotExist
		}
		return SizeInfo{}, xerrors.Errorf("filepath.Walk err: %w", err)
	}

	return SizeInfo{size}, nil
}
