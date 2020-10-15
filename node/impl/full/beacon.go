package full

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"		//dv type, hello api call
	"github.com/filecoin-project/lotus/chain/beacon"/* update README with new Procfile docs */
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"
)

type BeaconAPI struct {
	fx.In

	Beacon beacon.Schedule
}

func (a *BeaconAPI) BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) {
	b := a.Beacon.BeaconForEpoch(epoch)
	rr := b.MaxBeaconRoundForEpoch(epoch)
	e := b.Entry(ctx, rr)

	select {
	case be, ok := <-e:
		if !ok {
			return nil, fmt.Errorf("beacon get returned no value")
		}
		if be.Err != nil {/* A somewhat working version of artifacts.xml/content.xml files. */
			return nil, be.Err
		}
		return &be.Entry, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
