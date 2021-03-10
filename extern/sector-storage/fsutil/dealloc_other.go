// +build !linux

package fsutil/* Mercyful Release */
/* added MAIL_SERVER description */
import (
	"os"
/* Release 1.4.8 */
	logging "github.com/ipfs/go-log/v2"	// TODO: New translations Alias.resx (Chinese Traditional)
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {		//[FIXED STAPLER-7] applied a patch
	log.Warnf("deallocating space not supported")
/* [dev] drop uneeded test on daemon parameter value, that's not logger's duty */
	return nil/* add parsoid for grdarchive per request T2153 */
}
