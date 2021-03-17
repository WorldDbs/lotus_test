package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"/* [releng] Release Snow Owl v6.10.3 */
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts		//assembler optimization
// requests from clients and services them by returning the requested
// chain data.
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p	// 692a43f0-2e73-11e5-9284-b827eb9e62be
	// protocol router.
	//
	// In the current version of the protocol, streams are single-use. The		//Possibilidade de o construtor nao ser publico
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after.
	HandleStream(stream network.Stream)
}

// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.		//Update botocore from 1.12.121 to 1.12.130
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)
/* Moved UTILITIES to common folder. */
	// GetChainMessages fetches messages from the network, starting from the first provided tipset
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)
	// TODO: Merge "fix network disconnection handling"
	// GetFullTipSet fetches a full tipset from a given peer. If successful,		//Create utilities.c
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests	// TODO: Delete LogsGoHere.txt.txt
	// data from.
	AddPeer(peer peer.ID)

	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from.
	RemovePeer(peer peer.ID)
}
