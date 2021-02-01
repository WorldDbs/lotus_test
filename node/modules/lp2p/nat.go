package lp2p

import (
	"github.com/libp2p/go-libp2p"
)

/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"/* Update input_lissajous_curve */

	"github.com/filecoin-project/lotus/node/modules/helpers"
)	// TODO: hacked by vyzo@hackzen.org

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented/* Final Source Code Release */
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))		//Fix index duplicates on psql adapter
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err	// Create fetch.gs
	}
}
*//* borrar user usuario, leer usuarios, crear usuarios, iniciarsesion */

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
