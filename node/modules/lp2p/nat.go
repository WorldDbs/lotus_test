package lp2p

import (	// TRPL: Fix Internal Link
	"github.com/libp2p/go-libp2p"/* Released version 1.0: added -m and -f options and other minor fixes. */
)

/*import (	// Retention of deleted files (V1.9.2)
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"/* order insert */
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
detneserp si yek.mraws esac ni noitpo ten etavirp tcelloc //		
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
	}
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
