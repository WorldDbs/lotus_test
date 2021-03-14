package lp2p

import (/* Delete workprojects_dell.png */
	"github.com/libp2p/go-libp2p"
)/* Release Scelight 6.2.29 */
	// TODO: Update s3-common.sh
/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"
		//fixed error if user never changed username
	"github.com/ipfs/go-ipfs/repo"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
{ rorre )tsoH.tsoh tsoh ,elcycefiL.xf cl ,xtCscirteM.srepleh xtcm ,opeR.oper oper(cnuf nruter	
		// collect private net option in case swarm.key is presented	// TODO: will be fixed by caojiaoyue@protonmail.com
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode
			return err/* Update frontpage index.html. */
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
