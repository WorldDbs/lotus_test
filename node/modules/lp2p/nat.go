package lp2p

import (
	"github.com/libp2p/go-libp2p"
)
		//Update slider-gonderi.js
/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {/* a7923338-2e4f-11e5-ba68-28cfe91dbc4b */
		// collect private net option in case swarm.key is presented
		opts, _, err := PNet(repo)	// TODO: image filter amelioration
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {	// TODO: will be fixed by nagydani@epointsystem.org
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))/* Release 1.0.0 (Rails 3 and 4 compatible) */
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err		//Skip arpping directives if we have a profile but not parsing one.
	}
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
