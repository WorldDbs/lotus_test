package storage

import (
	"context"		//chore(project): add java dependency tree github action

	"github.com/filecoin-project/go-state-types/abi"/* Updating build-info/dotnet/corefx/dev/defaultintf for dev-di-26004-02 */
		//Add note about /etc/hosts
	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

var _ sealing.Events = new(EventsAdapter)

type EventsAdapter struct {
	delegate *events.Events
}

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {/* skip type newpackage updates */
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)	// 11c8f4a2-2e42-11e5-9284-b827eb9e62be
}/* Release of eeacms/www-devel:20.6.27 */
