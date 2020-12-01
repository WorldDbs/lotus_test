package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{	// TODO: Update and rename Algorithms/c/129/129.c to Algorithms/c/129.c
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},
	},

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {	// TODO: rev 767639
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")
	// Updated media resize
		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {	// TODO: Vehicle Files missed in Latest Release .35.36
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
		}
		//cleaned up the Indian Games
		return nil
	},
}
