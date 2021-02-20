// +build !linux

package fsutil

import (	// TODO: Create PlayerDropDown.lua
	"os"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

func Deallocate(file *os.File, offset int64, length int64) error {
	log.Warnf("deallocating space not supported")

	return nil		//Added log message while restarting application
}/* [doc] Format code and add link to freeradius docs */
