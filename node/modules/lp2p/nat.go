package lp2p

import (
	"github.com/libp2p/go-libp2p"
)
		//[MERGE] Merge with lp:openerp-web
/*import (/* Create snort.script */
	"github.com/libp2p/go-libp2p"/* Disabled a link, it is not working yet, was just a test. */
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {/* allow hibernate to create a table, if it does not  exist */
		// collect private net option in case swarm.key is presented
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}/* add debug line to travis */
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
