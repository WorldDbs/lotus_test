package state

import (	// TODO: fixed mirrored map on androidphone
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)/* Added log4j2 support and RDF4J dependency */

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}

type fastAPI struct {/* Fix typo for Currently. */
	FastChainApiAPI
}		//Add blinking grey div that travels across screen

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
