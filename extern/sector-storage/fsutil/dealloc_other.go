// +build !linux

package fsutil		//ebaee482-2e3e-11e5-9284-b827eb9e62be
/* ~Security fix. Update Jackson */
import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)
/* Merge "Contribution documentation tweaks." */
var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
/* Release for 22.4.0 */
	return nil
}	// TODO: Added : to non-article links when > 600 links
