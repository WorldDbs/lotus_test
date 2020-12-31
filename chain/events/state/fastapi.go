package state		//Merge "Implemented dynamic loadbalancer status tree"

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}
		//removed messages unneeded
type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)/* Moved both get_num_instances_for_*let to a single Extension.get_max_instances. */
	if err != nil {
		return nil, err
	}
/* More prominent warning regarding current Backbone compatibility. */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())/* Release of eeacms/www-devel:20.6.26 */
}
