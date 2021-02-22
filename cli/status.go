package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{
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
		defer closer()
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)/* FieldsAndGetters moved to DurianDebug. */
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)/* Release '0.2~ppa7~loms~lucid'. */
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)/* 6f09c322-2e5f-11e5-9284-b827eb9e62be */
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)
	// TODO: hacked by jon@atack.com
		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {/* Release areca-7.2.6 */
				ok100 = "[OK]"
			} else {/* Release 1.0.3 for Bukkit 1.5.2-R0.1 and ByteCart 1.5.0 */
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {/* Invoke compile tasks and force dependencies download running the `info` command */
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)	// TODO: will be fixed by hello@brooklynzelenka.com
		}

		return nil
	},
}
