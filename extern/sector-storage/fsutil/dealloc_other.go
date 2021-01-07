// +build !linux

package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"		//adding a note about the nightly binaries
)

var log = logging.Logger("fsutil")/* Delete logo.py */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")/* Add suspend confirm dialog */

	return nil
}
