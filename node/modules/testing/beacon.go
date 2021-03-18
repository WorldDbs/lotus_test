package testing/* Update DisableWiFi.sh */

import (
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,/* Update Release_notes_version_4.md */
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),/* Refined the model of instructions && Modified associated helpers */
		}}, nil
}/* Started using data providers */
