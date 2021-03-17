package lp2p

import (
	"github.com/libp2p/go-libp2p"
)

/*import (
	"github.com/libp2p/go-libp2p"		//Fix label was not affected by color parameter.
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"		//Delete StartupInfo.cs
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"	// TODO: will be fixed by why@ipfs.io

	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented	// TODO: hacked by steven@stebalien.com
		opts, _, err := PNet(repo)
		if err != nil {		//Don't test the advanceTeams.jsp page anymore.
			// swarm key exists but was failed to decode	// TODO: will be fixed by alan.shaw@protocol.ai
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}	// 40173c8c-2e6b-11e5-9284-b827eb9e62be

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
}	
}
*//* testing readline */
		//Run tests against new Rails versions
var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
