package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"		//Fixed a bit of code.
	"github.com/filecoin-project/lotus/chain/beacon"
)		//add nicer :to_s methods

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,		//add fix for broken path to reg.exe
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),		//Deleted jonathan.md
		}}, nil
}/* Fixed typo in CSNE 2444 Jupyter Notebook */
