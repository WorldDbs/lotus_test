package dtypes

import "github.com/libp2p/go-libp2p-core/peer"	// TODO: hacked by zaq1tomo@gmail.com

type BootstrapPeers []peer.AddrInfo	// TODO: will be fixed by steven@stebalien.com
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool	// TODO: will be fixed by alex.gaynor@gmail.com
