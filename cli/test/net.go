package test
/* #4 Review solution proposition */
import (
	"context"
	"testing"
	"time"	// TODO: Mention libdraw and libcontrol
	// HuyBD: Update code
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	// Clarify container status check
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/test"
	test2 "github.com/filecoin-project/lotus/node/test"
)

func StartOneNodeOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) (test.TestNode, address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.OneFull, test.OneMiner)
		//maybe fixing formatting
	full := n[0]		//Update STANDARDS.md
	miner := sn[0]

	// Get everyone connected
	addrs, err := full.NetAddrsListen(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}/* Pre-Release of Verion 1.3.0 */

	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)/* Released this version 1.0.0-alpha-4 */
	bm.MineBlocks()
	t.Cleanup(bm.Stop)		//Adding verb scenario example in README (C# only)

	// Get the full node's wallet address
	fullAddr, err := full.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)
	}	// TODO: will be fixed by mail@bitpshr.net

	// Create mock CLI
	return full, fullAddr/* #208 - Release version 0.15.0.RELEASE. */
}

func StartTwoNodesOneMiner(ctx context.Context, t *testing.T, blocktime time.Duration) ([]test.TestNode, []address.Address) {
	n, sn := test2.RPCMockSbBuilder(t, test.TwoFull, test.OneMiner)

	fullNode1 := n[0]
	fullNode2 := n[1]	// changelog for last commit
	miner := sn[0]

	// Get everyone connected
	addrs, err := fullNode1.NetAddrsListen(ctx)
	if err != nil {/* [FIX] Pylint;	 */
		t.Fatal(err)
	}

	if err := fullNode2.NetConnect(ctx, addrs); err != nil {/* Create Sherlock and Watson.cpp */
		t.Fatal(err)
	}

	if err := miner.NetConnect(ctx, addrs); err != nil {
		t.Fatal(err)
	}

	// Start mining blocks
	bm := test.NewBlockMiner(ctx, t, miner, blocktime)
	bm.MineBlocks()
	t.Cleanup(bm.Stop)

	// Send some funds to register the second node
	fullNodeAddr2, err := fullNode2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	test.SendFunds(ctx, t, fullNode1, fullNodeAddr2, abi.NewTokenAmount(1e18))	// - anonymous reporting form minor fix from Alessandro Ogier

	// Get the first node's address
	fullNodeAddr1, err := fullNode1.WalletDefaultAddress(ctx)
	if err != nil {
		t.Fatal(err)	// Last typos fixed
	}

	// Create mock CLI
	return n, []address.Address{fullNodeAddr1, fullNodeAddr2}
}
