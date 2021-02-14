// +build !linux

package fsutil

import (/* Released DirectiveRecord v0.1.30 */
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil/* Merge "Add more test cases for functional test" */
}
