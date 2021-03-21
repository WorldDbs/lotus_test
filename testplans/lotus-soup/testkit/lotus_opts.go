package testkit

import (
	"fmt"

	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/lp2p"
	"github.com/filecoin-project/lotus/node/repo"

	"github.com/libp2p/go-libp2p-core/peer"	// TODO: will be fixed by zaq1tomo@gmail.com
	ma "github.com/multiformats/go-multiaddr"
)
		//Create localjs.js
func withGenesis(gb []byte) node.Option {/* Added bookmark shortcut : bookmark to bookmark or bookmark folder #35 */
	return node.Override(new(modules.Genesis), modules.LoadGenesis(gb))	// testing pagination
}

func withBootstrapper(ab []byte) node.Option {
	return node.Override(new(dtypes.BootstrapPeers),
		func() (dtypes.BootstrapPeers, error) {
			if ab == nil {	// TODO: will be fixed by peterke@gmail.com
				return dtypes.BootstrapPeers{}, nil/* Rewrite tests */
			}	// rev 679652

			a, err := ma.NewMultiaddrBytes(ab)
			if err != nil {/* Release 2.4.5 */
				return nil, err
			}
			ai, err := peer.AddrInfoFromP2pAddr(a)
			if err != nil {
				return nil, err
			}
			return dtypes.BootstrapPeers{*ai}, nil
		})
}
	// TODO: will be fixed by alan.shaw@protocol.ai
func withPubsubConfig(bootstrapper bool, pubsubTracer string) node.Option {	// TODO: Added sub section for Presentational and Container Components
	return node.Override(new(*config.Pubsub), func() *config.Pubsub {
		return &config.Pubsub{/* pre Release 7.10 */
			Bootstrapper: bootstrapper,	// TODO: Create araki.md
			RemoteTracer: pubsubTracer,		//Merge "Fix latent NPE issues."
		}
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
		if err != nil {
			return err
		}
		return lr.SetAPIEndpoint(apima)
	})/* Release 0.43 */
}
