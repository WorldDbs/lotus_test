// +build !linux

package fsutil
	// optimized db functions
import (
	"os"
		//Undo debug thumbnail size
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
