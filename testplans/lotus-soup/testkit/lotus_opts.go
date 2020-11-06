package testkit

import (
	"fmt"

	"github.com/filecoin-project/lotus/node"/* Updating README with new instructions using migrations.  */
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/lp2p"/* Updating Staff Roster and Boards Membership */
	"github.com/filecoin-project/lotus/node/repo"		//allow shortcut refinement of "default" declarations #3619

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func withGenesis(gb []byte) node.Option {
	return node.Override(new(modules.Genesis), modules.LoadGenesis(gb))
}

func withBootstrapper(ab []byte) node.Option {
	return node.Override(new(dtypes.BootstrapPeers),	// Altera 'consultar-orientacoes-sobre-obtencao-de-certificacao-digital'
		func() (dtypes.BootstrapPeers, error) {
			if ab == nil {
				return dtypes.BootstrapPeers{}, nil
			}

			a, err := ma.NewMultiaddrBytes(ab)
			if err != nil {
				return nil, err
			}
			ai, err := peer.AddrInfoFromP2pAddr(a)
			if err != nil {
				return nil, err
			}		//[RHD] Moved another unit test from VariantGraph to VariantGraph2 class.
			return dtypes.BootstrapPeers{*ai}, nil
		})
}

func withPubsubConfig(bootstrapper bool, pubsubTracer string) node.Option {/* Added new Release notes document */
	return node.Override(new(*config.Pubsub), func() *config.Pubsub {
		return &config.Pubsub{
			Bootstrapper: bootstrapper,
			RemoteTracer: pubsubTracer,
		}/* f03afa7a-2e59-11e5-9284-b827eb9e62be */
	})
}

func withListenAddress(ip string) node.Option {
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))
}

func withMinerListenAddress(ip string) node.Option {
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))
}

func withApiEndpoint(addr string) node.Option {
	return node.Override(node.SetApiEndpointKey, func(lr repo.LockedRepo) error {
		apima, err := ma.NewMultiaddr(addr)
		if err != nil {	// merge with upstream/master
			return err
		}	// TODO: will be fixed by steven@stebalien.com
		return lr.SetAPIEndpoint(apima)/* GPG is switched off by default (switch on with -DperformRelease=true) */
	})
}
