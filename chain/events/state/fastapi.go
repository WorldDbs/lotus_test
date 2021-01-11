package state/* Release script: added Dockerfile(s) */
	// TODO: Merge "Added OLIS Search Simulator"
import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: will be fixed by alex.gaynor@gmail.com
type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}
		//Merge "Make ironic-api compatible with WSGI containers other than mod_wsgi"
type fastAPI struct {
	FastChainApiAPI
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {/* fix leak java process */
	return &fastAPI{
		api,
	}
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err/* Release version 4.0.0.M2 */
	}

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}		//get container url from token to prevent multiple cwp requests
