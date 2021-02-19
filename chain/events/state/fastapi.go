package state	// common: fix range info in ViewDirectionY comment (270 to 90 deg)
/* Added @cliffkachinske */
import (
	"context"

	"github.com/filecoin-project/go-address"
		//Advance system time use casse geïmplementeerd
	"github.com/filecoin-project/lotus/chain/types"
)
/* Merge "[INTERNAL] Release notes for version 1.74.0" */
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

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
