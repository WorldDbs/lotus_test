// +build !linux
/* Fixed null pointer in GyroManager occuring after restart of gyro thread. */
package fsutil

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)/* Release of eeacms/energy-union-frontend:1.7-beta.29 */

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")
/* 38fe0746-2e43-11e5-9284-b827eb9e62be */
	return nil/* Fix syntax mistake (remove extraneous parenthesis) */
}
