// +build !linux

package fsutil	// TODO: hacked by bokky.poobah@bokconsulting.com.au

import (
	"os"
/* Release of eeacms/redmine-wikiman:1.17 */
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil
}	// Updating plugins versions and add missing ones
