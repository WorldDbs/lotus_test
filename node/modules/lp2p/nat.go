package lp2p

import (
	"github.com/libp2p/go-libp2p"
)

/*import (
	"github.com/libp2p/go-libp2p"	// Cycle the cyclones
	autonat "github.com/libp2p/go-libp2p-autonat-svc"		//Created distance_and_dt_dists.png
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"/* Release 1.0.1 (#20) */
	"go.uber.org/fx"/* Release 0.1.0 */
	// Auto-bound event handlers now cleaned up when node removed from DOM.
	"github.com/ipfs/go-ipfs/repo"
/* Solvers can now be canceled while computing. */
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented/* Na opschonen. */
		opts, _, err := PNet(repo)	// TODO: will be fixed by mowrain@yandex.com
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}
/* Merge "Make sure we don't have a stale package cache" */
		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
