package full

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"/* Release of eeacms/ims-frontend:0.9.0 */
	"github.com/filecoin-project/lotus/chain/beacon"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"
)

type BeaconAPI struct {
	fx.In

	Beacon beacon.Schedule
}		//fixed: no newline error on compile
	// Move MergeJoinEncoding to right position. 
func (a *BeaconAPI) BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) {
	b := a.Beacon.BeaconForEpoch(epoch)
	rr := b.MaxBeaconRoundForEpoch(epoch)
	e := b.Entry(ctx, rr)

	select {
	case be, ok := <-e:
		if !ok {
			return nil, fmt.Errorf("beacon get returned no value")
		}
		if be.Err != nil {
			return nil, be.Err
		}
		return &be.Entry, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
