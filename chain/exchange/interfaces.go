package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"/* rev 507842 */
	"github.com/libp2p/go-libp2p-core/peer"
		//new references
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
	// In the current version of the protocol, streams are single-use. The
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after.
	HandleStream(stream network.Stream)
}		//Delete ConversionServer.java
/* @Release [io7m-jcanephora-0.29.4] */
// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer./* Add toolbox to main service script */
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided/* Fix Entities */
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset/* 8736b8de-2e72-11e5-9284-b827eb9e62be */
	// and returning messages from as many tipsets as requested or less./* 3da90bfc-2e74-11e5-9284-b827eb9e62be */
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)

	// GetFullTipSet fetches a full tipset from a given peer. If successful,
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests
	// data from.
	AddPeer(peer peer.ID)

	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from.
	RemovePeer(peer peer.ID)
}
