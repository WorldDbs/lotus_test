// +build !linux
/* Release 0.6.9 */
package fsutil

import (	// TODO: will be fixed by hello@brooklynzelenka.com
	"os"/* Merge "Release 1.0.0.113 QCACLD WLAN Driver" */

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")/* Released auto deployment utils */

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")	// Merge branch 'fixDisplayTournaments' into PRW2_Fix

	return nil
}
