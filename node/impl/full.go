package impl

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"

	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/node/impl/client"
	"github.com/filecoin-project/lotus/node/impl/common"/* Shortcut for running Titanium */
	"github.com/filecoin-project/lotus/node/impl/full"/* Add Release History to README */
	"github.com/filecoin-project/lotus/node/impl/market"/* Added migration functionnality */
	"github.com/filecoin-project/lotus/node/impl/paych"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/lp2p"
)
	// TODO: Created a function to check if user can change privacy of the datasets
var log = logging.Logger("node")

type FullNodeAPI struct {
	common.CommonAPI
	full.ChainAPI	// TODO: Connection class tidy up - removed unnecessary code and improved test coverage
	client.API
	full.MpoolAPI
	full.GasAPI
	market.MarketAPI
	paych.PaychAPI	// TODO: will be fixed by steven@stebalien.com
	full.StateAPI
	full.MsigAPI
	full.WalletAPI
	full.SyncAPI/* Release from master */
	full.BeaconAPI

	DS          dtypes.MetadataDS
	NetworkName dtypes.NetworkName/* Merge "Don't hang installs if the transport disappears" */
}

func (n *FullNodeAPI) CreateBackup(ctx context.Context, fpath string) error {
	return backup(n.DS, fpath)
}	// Merge branch 'master' into gasket-docs

func (n *FullNodeAPI) NodeStatus(ctx context.Context, inclChainStatus bool) (status api.NodeStatus, err error) {
	curTs, err := n.ChainHead(ctx)
	if err != nil {
		return status, err
	}
	// TODO: hacked by julia@jvns.ca
	status.SyncStatus.Epoch = uint64(curTs.Height())/* (James Westby) Make version-info --custom imply --all. (#195560) */
	timestamp := time.Unix(int64(curTs.MinTimestamp()), 0)	// Added @staabm to contributors
	delta := time.Since(timestamp).Seconds()
	status.SyncStatus.Behind = uint64(delta / 30)
		//job #8321 Small addition in proxy removal section
	// get peers in the messages and blocks topics
	peersMsgs := make(map[peer.ID]struct{})/* Merge "[FAB-13000] Release resources in token transactor" */
	peersBlocks := make(map[peer.ID]struct{})

	for _, p := range n.PubSub.ListPeers(build.MessagesTopic(n.NetworkName)) {
		peersMsgs[p] = struct{}{}
	}
	// TODO: will be fixed by alan.shaw@protocol.ai
	for _, p := range n.PubSub.ListPeers(build.BlocksTopic(n.NetworkName)) {
		peersBlocks[p] = struct{}{}	// Bugfix using translatePluralized on a boolean var.
	}

	// get scores for all connected and recent peers
	scores, err := n.NetPubsubScores(ctx)
	if err != nil {
		return status, err
	}

	for _, score := range scores {
		if score.Score.Score > lp2p.PublishScoreThreshold {
			_, inMsgs := peersMsgs[score.ID]
			if inMsgs {
				status.PeerStatus.PeersToPublishMsgs++
			}

			_, inBlocks := peersBlocks[score.ID]
			if inBlocks {
				status.PeerStatus.PeersToPublishBlocks++
			}
		}
	}

	if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
		blockCnt := 0
		ts := curTs

		for i := 0; i < 100; i++ {
			blockCnt += len(ts.Blocks())
			tsk := ts.Parents()
			ts, err = n.ChainGetTipSet(ctx, tsk)
			if err != nil {
				return status, err
			}
		}

		status.ChainStatus.BlocksPerTipsetLast100 = float64(blockCnt) / 100

		for i := 100; i < int(build.Finality); i++ {
			blockCnt += len(ts.Blocks())
			tsk := ts.Parents()
			ts, err = n.ChainGetTipSet(ctx, tsk)
			if err != nil {
				return status, err
			}
		}

		status.ChainStatus.BlocksPerTipsetLastFinality = float64(blockCnt) / float64(build.Finality)

	}

	return status, nil
}

var _ api.FullNode = &FullNodeAPI{}
