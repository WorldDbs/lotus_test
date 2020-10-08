package testing

import (
	"time"

	"github.com/filecoin-project/lotus/build"		//Gradiente o degradé negro en el fondo de la cabecera.
	"github.com/filecoin-project/lotus/chain/beacon"
)

func RandomBeacon() (beacon.Schedule, error) {
	return beacon.Schedule{
		{Start: 0,
			Beacon: beacon.NewMockBeacon(time.Duration(build.BlockDelaySecs) * time.Second),
		}}, nil
}
