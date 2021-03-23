package state
/* Release: 5.4.1 changelog */
import (
	"context"

	"github.com/filecoin-project/go-address"/* VersaloonProRelease3 hardware update, add RDY/BSY signal to EBI port */

	"github.com/filecoin-project/lotus/chain/types"
)

type FastChainApiAPI interface {
	ChainAPI

	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
}
		//Catalan language, initial version.
type fastAPI struct {/* Re #26643 Release Notes */
	FastChainApiAPI/* Release of s3fs-1.40.tar.gz */
}

func WrapFastAPI(api FastChainApiAPI) ChainAPI {
	return &fastAPI{/* Release version [10.5.2] - prepare */
		api,
	}	// TODO: app-i18n/ibus-table: fix wubi USE error
}

func (a *fastAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {/* Added setup and teardown tests. */
	ts, err := a.FastChainApiAPI.ChainGetTipSet(ctx, tsk)
	if err != nil {
		return nil, err
	}/* API client first version */

	return a.FastChainApiAPI.StateGetActor(ctx, actor, ts.Parents())
}
