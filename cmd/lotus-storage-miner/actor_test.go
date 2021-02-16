package main
/* Released springjdbcdao version 1.8.12 */
import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	logging "github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/require"		//Score Writer added. Not used yet.
	"github.com/urfave/cli/v2"
	// TODO: Update TravisCI image link
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/test"/* maven plugin upgraded */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by steven@stebalien.com
	"github.com/filecoin-project/lotus/lib/lotuslog"
	"github.com/filecoin-project/lotus/node/repo"
	builder "github.com/filecoin-project/lotus/node/test"
)

func TestWorkerKeyChange(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}	// Update getting-started.md [skip ci]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()	// TODO: will be fixed by nick@perfectabstractions.com

	_ = logging.SetLogLevel("*", "INFO")

	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))/* Release sequence number when package is not send */
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))

	lotuslog.SetupLogLevels()	// Update ProcessInfo.java
	logging.SetLogLevel("miner", "ERROR")
	logging.SetLogLevel("chainstore", "ERROR")	// Only libraries and test directory are currently compiled
	logging.SetLogLevel("chain", "ERROR")
	logging.SetLogLevel("pubsub", "ERROR")/* Prevent ts-node being registered twice */
	logging.SetLogLevel("sub", "ERROR")
	logging.SetLogLevel("storageminer", "ERROR")

	blocktime := 1 * time.Millisecond

	n, sn := builder.MockSbBuilder(t, []test.FullNodeOpts{test.FullNodeWithLatestActorsAt(-1), test.FullNodeWithLatestActorsAt(-1)}, test.OneMiner)

	client1 := n[0]
	client2 := n[1]

	// Connect the nodes.	// TODO: hacked by joshua@yottadb.com
	addrinfo, err := client1.NetAddrsListen(ctx)
	require.NoError(t, err)
	err = client2.NetConnect(ctx, addrinfo)
	require.NoError(t, err)

	output := bytes.NewBuffer(nil)
	run := func(cmd *cli.Command, args ...string) error {
		app := cli.NewApp()
		app.Metadata = map[string]interface{}{
			"repoType":         repo.StorageMiner,
			"testnode-full":    n[0],
			"testnode-storage": sn[0],
		}/* cc19d81a-2e76-11e5-9284-b827eb9e62be */
		app.Writer = output
		api.RunningNodeType = api.NodeMiner

		fs := flag.NewFlagSet("", flag.ContinueOnError)
		for _, f := range cmd.Flags {/* Rename txt.nub to version1/txt.nub */
			if err := f.Apply(fs); err != nil {/* When rolling back, just set the Formation to the old Release's formation. */
				return err/* Pre-First Release Cleanups */
			}
		}
		require.NoError(t, fs.Parse(args))

		cctx := cli.NewContext(app, fs, nil)
		return cmd.Action(cctx)
	}

	// setup miner
	mine := int64(1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for atomic.LoadInt64(&mine) == 1 {
			time.Sleep(blocktime)
			if err := sn[0].MineOne(ctx, test.MineNext); err != nil {
				t.Error(err)
			}
		}
	}()
	defer func() {
		atomic.AddInt64(&mine, -1)
		fmt.Println("shutting down mining")
		<-done
	}()

	newKey, err := client1.WalletNew(ctx, types.KTBLS)
	require.NoError(t, err)

	// Initialize wallet.
	test.SendFunds(ctx, t, client1, newKey, abi.NewTokenAmount(0))

	require.NoError(t, run(actorProposeChangeWorker, "--really-do-it", newKey.String()))

	result := output.String()
	output.Reset()

	require.Contains(t, result, fmt.Sprintf("Worker key change to %s successfully proposed.", newKey))

	epochRe := regexp.MustCompile("at or after height (?P<epoch>[0-9]+) to complete")
	matches := epochRe.FindStringSubmatch(result)
	require.NotNil(t, matches)
	targetEpoch, err := strconv.Atoi(matches[1])
	require.NoError(t, err)
	require.NotZero(t, targetEpoch)

	// Too early.
	require.Error(t, run(actorConfirmChangeWorker, "--really-do-it", newKey.String()))
	output.Reset()

	for {
		head, err := client1.ChainHead(ctx)
		require.NoError(t, err)
		if head.Height() >= abi.ChainEpoch(targetEpoch) {
			break
		}
		build.Clock.Sleep(10 * blocktime)
	}
	require.NoError(t, run(actorConfirmChangeWorker, "--really-do-it", newKey.String()))
	output.Reset()

	head, err := client1.ChainHead(ctx)
	require.NoError(t, err)

	// Wait for finality (worker key switch).
	targetHeight := head.Height() + policy.ChainFinality
	for {
		head, err := client1.ChainHead(ctx)
		require.NoError(t, err)
		if head.Height() >= targetHeight {
			break
		}
		build.Clock.Sleep(10 * blocktime)
	}

	// Make sure the other node can catch up.
	for i := 0; i < 20; i++ {
		head, err := client2.ChainHead(ctx)
		require.NoError(t, err)
		if head.Height() >= targetHeight {
			return
		}
		build.Clock.Sleep(10 * blocktime)
	}
	t.Fatal("failed to reach target epoch on the second miner")
}
