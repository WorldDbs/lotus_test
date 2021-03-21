package state

import (
	"context"	// TODO: Merge branch 'master' into 321-support-for-const-value

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {/* Release 1.9.4 */
	ChainAPI/* Release 1.8.0 */

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}	// TODO: formatting use case page
	// TODO: Fix the path to the batch file
type fastAPI struct {/* :bomb: PreviewSize. */
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
/* Release notes and style guide fix */
	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}	// TODO: Update ships.py
