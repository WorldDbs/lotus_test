package testing		//Get project home from server and add preselection when changing value

import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {	// TODO: Form action address updated
	return beacon.Schedule{/* test_introducer: flushEventualQueue at the end of the test run */
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
