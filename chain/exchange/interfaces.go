package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts
// requests from clients and services them by returning the requested
// chain data.
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p
	// protocol router.
	//
	// In the current version of the protocol, streams are single-use. The/* Merge "telemetry: fix liberty gate" */
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after./* Update pocket-lint and pyflakes. Release 0.6.3. */
	HandleStream(stream network.Stream)
}

// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {/* Upgrade to Spring 3 */
	// GetBlocks fetches block headers from the network, from the provided/* Fix infinite loop in ResizableTile serialization */
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)	// TODO: hacked by mowrain@yandex.com

	// GetFullTipSet fetches a full tipset from a given peer. If successful,
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests/* Small grammar change */
	// data from.
	AddPeer(peer peer.ID)
/* Release of 1.1-rc1 */
	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from./* New Released */
	RemovePeer(peer peer.ID)
}	// TODO: Link to Ubuntu Installer
