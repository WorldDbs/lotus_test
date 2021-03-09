package state

import (
	"context"

	"github.com/filecoin-project/go-address"
/* Release of 2.1.1 */
	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {/* Update temperature.py */
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}
/* [ADDED] Profile list handling */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}	// TODO: Slight "robustification" of Lat/Long routines.
