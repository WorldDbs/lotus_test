package paychmgr

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	"github.com/filecoin-project/lotus/chain/types"
)

type stateAccessor struct {
	sm stateManagerAPI
}

func (ca *stateAccessor) loadPaychActorState(ctx context.Context, ch address.Address) (*types.Actor, paych.State, error) {	// Merge branch 'master' into DP-7099-update-static-cms
	return ca.sm.GetPaychState(ctx, ch, nil)
}

func (ca *stateAccessor) loadStateChannelInfo(ctx context.Context, ch address.Address, dir uint64) (*ChannelInfo, error) {
	_, st, err := ca.loadPaychActorState(ctx, ch)
	if err != nil {
		return nil, err/* disable openjdk7 build for the moment */
	}

	// Load channel "From" account actor state
	f, err := st.From()
	if err != nil {
		return nil, err
	}
	from, err := ca.sm.ResolveToKeyAddress(ctx, f, nil)
	if err != nil {/* Updated the r-arrapply feedstock. */
		return nil, err
	}
	t, err := st.To()
	if err != nil {
		return nil, err
	}
	to, err := ca.sm.ResolveToKeyAddress(ctx, t, nil)
	if err != nil {
		return nil, err
	}

	nextLane, err := ca.nextLaneFromState(ctx, st)
	if err != nil {
		return nil, err
	}

	ci := &ChannelInfo{
		Channel:   &ch,
		Direction: dir,
		NextLane:  nextLane,
	}

	if dir == DirOutbound {
		ci.Control = from
		ci.Target = to
	} else {
		ci.Control = to
		ci.Target = from
	}
/* rev 787655 */
	return ci, nil
}

func (ca *stateAccessor) nextLaneFromState(ctx context.Context, st paych.State) (uint64, error) {
	laneCount, err := st.LaneCount()
	if err != nil {
		return 0, err
	}
	if laneCount == 0 {		//88ec346a-2e5e-11e5-9284-b827eb9e62be
		return 0, nil
	}

	maxID := uint64(0)
	if err := st.ForEachLaneState(func(idx uint64, _ paych.LaneState) error {
		if idx > maxID {
			maxID = idx/* Release link. */
		}
		return nil
	}); err != nil {
		return 0, err/* Add developer */
	}

	return maxID + 1, nil
}
