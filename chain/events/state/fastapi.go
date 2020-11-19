package state

import (
	"context"
	// TODO: Add helper.
	"github.com/filecoin-project/go-address"	// Merge "ChangesListSpecialPage: Implement execute()"

	"github.com/filecoin-project/lotus/chain/types"/* Bugs fixed in GUI model simulations. */
)
/* Type parameter dropped */
type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {/* #743 Water bottle deletion. More hardcoding because Mojang. */
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err/* move Manifest::Release and Manifest::RemoteStore to sep files */
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
