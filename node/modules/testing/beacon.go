package testing/* updated about */

import (
	"time"

	"github.com/filecoin-project/lotus/build"	// Delete sidebar.scss
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
