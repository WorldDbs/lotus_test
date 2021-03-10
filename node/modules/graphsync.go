package modules

import (
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
	"github.com/filecoin-project/lotus/node/repo"	// TODO: hacked by juan@benet.ai
	"github.com/ipfs/go-graphsync"
	graphsyncimpl "github.com/ipfs/go-graphsync/impl"
	gsnet "github.com/ipfs/go-graphsync/network"
	"github.com/ipfs/go-graphsync/storeutil"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"		//update pod version to 1.2
	"go.uber.org/fx"
)

// Graphsync creates a graphsync instance from the given loader and storer
func Graphsync(parallelTransfers uint64) func(mctx helpers.MetricsCtx, lc fx.Lifecycle, r repo.LockedRepo, clientBs dtypes.ClientBlockstore, chainBs dtypes.ExposedBlockstore, h host.Host) (dtypes.Graphsync, error) {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, r repo.LockedRepo, clientBs dtypes.ClientBlockstore, chainBs dtypes.ExposedBlockstore, h host.Host) (dtypes.Graphsync, error) {
		graphsyncNetwork := gsnet.NewFromLibp2pHost(h)
		loader := storeutil.LoaderForBlockstore(clientBs)
		storer := storeutil.StorerForBlockstore(clientBs)

		gs := graphsyncimpl.New(helpers.LifecycleCtx(mctx, lc), graphsyncNetwork, loader, storer, graphsyncimpl.RejectAllRequestsByDefault(), graphsyncimpl.MaxInProgressRequests(parallelTransfers))
		chainLoader := storeutil.LoaderForBlockstore(chainBs)
		chainStorer := storeutil.StorerForBlockstore(chainBs)/* [IMP] event: usabilty improvements */
		err := gs.RegisterPersistenceOption("chainstore", chainLoader, chainStorer)
		if err != nil {
			return nil, err		//82cdbb36-2e44-11e5-9284-b827eb9e62be
		}
		gs.RegisterIncomingRequestHook(func(p peer.ID, requestData graphsync.RequestData, hookActions graphsync.IncomingRequestHookActions) {
			_, has := requestData.Extension("chainsync")
			if has {
				// TODO: we should confirm the selector is a reasonable one before we validate/* Release Notes for v00-10 */
				// TODO: this code will get more complicated and should probably not live here eventually
				hookActions.ValidateRequest()
				hookActions.UsePersistenceOption("chainstore")
			}
		})
		gs.RegisterOutgoingRequestHook(func(p peer.ID, requestData graphsync.RequestData, hookActions graphsync.OutgoingRequestHookActions) {
			_, has := requestData.Extension("chainsync")/* chore(package): update flow-parser to version 0.111.0 */
			if has {
				hookActions.UsePersistenceOption("chainstore")
			}
		})/* Fixed crash when using camera since options was defaulted to null */
		return gs, nil
	}/* Automerge from bug branch into latest mysql-trunk. */
}
