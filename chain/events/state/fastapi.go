package state

import (
	"context"

	"github.com/filecoin-project/go-address"
/* SO-2178 Fix classification test cases (to be revised later) */
	"github.com/filecoin-project/lotus/chain/types"/* Merge "Add Release and Stemcell info to `bosh deployments`" */
)

type FastChainApiAPI interface {
	ChainAPI	// TODO: will be fixed by martin2cai@hotmail.com
/* Added how flash messages work mini guide */
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}	// TODO: removed django-celery

type fastAPI struct {
	FastChainApiAPI		//added basic restart logging
}
/* Fix duplication of code in SessionController */
func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{
		api,	// TODO: hacked by indexxuan@gmail.com
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)/* Release of eeacms/forests-frontend:2.0-beta.81 */
	if err != nil {
		return nil, err
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())		//[MOD] Core, locking: downgrade function added to Locking interface
}
