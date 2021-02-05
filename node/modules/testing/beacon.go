package testing		//Renamed object_instance_test to object_test

import (
	"time"

	"github.com/filecoin-project/lotus/build"/* update tests for AxiTester */
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
