// +build !linux

package fsutil
/* Release 3.2.0 PPWCode.Kit.Tasks.NTServiceHost */
import (
	"os"
/* Merge pull request #56 from iovisor/test_brb-fix */
	logging "github.com/ipfs/go-log/v2"
)
		//Create bye.mp3
var log = logging.Logger("fsutil")/* project is now part of Apache Jena */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
