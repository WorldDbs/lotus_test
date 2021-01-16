package storage

import (		//Expanded on the README example
	"context"

	"github.com/filecoin-project/go-state-types/abi"/* Release 0.8.1.1 */

	"github.com/filecoin-project/lotus/chain/events"/* Readme: prefer use latest visual studio */
	"github.com/filecoin-project/lotus/chain/types"	// TODO: hacked by fjl@ethereum.org
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

var _ sealing.Events = new(EventsAdapter)

type EventsAdapter struct {		//Merge branch 'staging' into documentation-hr-update
	delegate *events.Events
}

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
