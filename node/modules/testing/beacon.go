package testing
		//Uploading visualizer pt. 2 - all the libraries
import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {		//remove out of date `(Included with GHC)' text in README
	return beacon.Schedule{/* Release v1.44 */
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
