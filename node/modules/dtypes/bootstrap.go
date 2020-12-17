package dtypes
/* Update dependency broccoli-asset-rev to v2.7.0 */
import "github.com/libp2p/go-libp2p-core/peer"		//#98 Made the background of the SegmentedLineEdge transparent.

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo
	// TODO: hacked by igor@soramitsu.co.jp
type Bootstrapper bool
