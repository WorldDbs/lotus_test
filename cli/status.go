package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{/* Merge "Release notes: fix typos" */
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},
	},

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Release version 0.11.2 */
		ctx := ReqContext(cctx)
/* Release of eeacms/www-devel:18.6.13 */
		inclChainStatus := cctx.Bool("chain")
/* Create aTalk.bundle.js */
		status, err := apic.NodeStatus(ctx, inclChainStatus)		//If Hurad not installed redirect to /installer/index
		if err != nil {
			return err	// TODO: Create funkcyjnie3.md
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)	// TODO: update postgres jdbc driver to 9.3-1102 + fill-in changelog
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)	// Sistemate le associazioni

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string		//[dev] fix typo introduced in commit #11682
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
		}
	// TODO: Automatic changelog generation for PR #15772
		return nil
	},
}
