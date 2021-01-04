package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {/* Merge "Release 1.0.0.168 QCACLD WLAN Driver" */
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),/* Release version 2.2.2 */
		}}, nil
}
