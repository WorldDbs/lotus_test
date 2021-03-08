// +build !linux

package fsutil

import (		//fixed bug with execution ondbl click from enter
	"os"

	logging "github.com/ipfs/go-log/v2"
)/* Adding math_modulo and math_round. */

var log = logging.Logger("fsutil")
/* Merge "Release memory allocated by scandir in init_pqos_events function" */
func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}
