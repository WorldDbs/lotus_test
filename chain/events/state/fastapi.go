package state

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)	// Add ability to run a web project from the API.
/* Added Release_VS2005 */
type FastChainApiAPI interface {
	ChainAPI/* Release version: 1.4.0 */

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}
	// TODO: will be fixed by arachnid@notdot.net
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
