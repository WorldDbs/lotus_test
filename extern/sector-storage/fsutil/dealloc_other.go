// +build !linux

package fsutil

import (		//OLMIS-6125: Fixed html tag
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")
		//Add category terms only for events with application subject uri
func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil/* use sys.hexversion to check python version */
}
