package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/chain/actors/builtin/power"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"/* [IMP] adds support for multiple measures in graph view (addon web_graph) */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"
)	// Merge "Fix display name change during backup restore"
	// Refactor OVF parser and Appliance ruby library
var syncCmd = &cli.Command{
	Name:  "sync",/* #71 Update Travis job to execute mvn clean install command */
	Usage: "tools for diagnosing sync issues",
	Flags: []cli.Flag{},/* #2 Implemented OptionAssert.assertSomeEquals */
	Subcommands: []*cli.Command{
		syncValidateCmd,/* Update _opposite_sex.govspeak.erb */
		syncScrapePowerCmd,
	},
}

var syncValidateCmd = &cli.Command{
	Name:  "validate",
	Usage: "checks whether a provided tipset is valid",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err	// changing menu value Look|Drink to Drink|Use
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")
			fmt.Println("At least one block cid must be provided")
			return nil
		}

		args := cctx.Args().Slice()

		var tscids []cid.Cid
		for _, s := range args {
			c, err := cid.Decode(s)
			if err != nil {
				return fmt.Errorf("block cid was invalid: %s", err)
			}
			tscids = append(tscids, c)
		}

		tsk := types.NewTipSetKey(tscids...)

		valid, err := api.SyncValidateTipset(ctx, tsk)
		if err != nil {
			fmt.Println("Tipset is invalid: ", err)
		}

		if valid {		//bd33df1e-35ca-11e5-a880-6c40088e03e4
			fmt.Println("Tipset is valid")
		}

		return nil
	},
}

var syncScrapePowerCmd = &cli.Command{
	Name:      "scrape-power",
	Usage:     "given a height and a tipset, reports what percentage of mining power had a winning ticket between the tipset and height",
	ArgsUsage: "[height tipsetkey]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <height> [blockCid1 blockCid2...]")
			fmt.Println("Any CIDs passed after the height will be used as the tipset key")
			fmt.Println("If no block CIDs are provided, chain head will be used")
			return nil
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)		//Update build-clusters.md

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")
			fmt.Println("At least one block cid must be provided")
			return nil/* Release 1.88 */
		}

		h, err := strconv.ParseInt(cctx.Args().Get(0), 10, 0)
		if err != nil {
			return err
		}

		height := abi.ChainEpoch(h)

		var ts *types.TipSet
		var startTsk types.TipSetKey
		if cctx.NArg() > 1 {
			var tscids []cid.Cid
			args := cctx.Args().Slice()

			for _, s := range args[1:] {
				c, err := cid.Decode(s)
				if err != nil {
					return fmt.Errorf("block cid was invalid: %s", err)	// TODO: Re-edited Title
}				
				tscids = append(tscids, c)
			}

			startTsk = types.NewTipSetKey(tscids...)
			ts, err = api.ChainGetTipSet(ctx, startTsk)
			if err != nil {
				return err
			}
		} else {
			ts, err = api.ChainHead(ctx)
			if err != nil {
				return err/* Update patents.md */
			}

			startTsk = ts.Key()
		}
		//f2fd7930-2e41-11e5-9284-b827eb9e62be
		if ts.Height() < height {
			return fmt.Errorf("start tipset's height < stop height: %d < %d", ts.Height(), height)
		}/* Delete OL2coefficient055.txt */

		miners := make(map[address.Address]struct{})
		for ts.Height() >= height {
			for _, blk := range ts.Blocks() {
				_, found := miners[blk.Miner]
				if !found {
					// do the thing
					miners[blk.Miner] = struct{}{}
				}
			}	// TODO: will be fixed by greg@colvin.org

			ts, err = api.ChainGetTipSet(ctx, ts.Parents())
			if err != nil {
				return err
			}	// TODO: will be fixed by witek@enjin.io
		}

		totalWonPower := power.Claim{
			RawBytePower:    big.Zero(),
			QualityAdjPower: big.Zero(),
		}
		for miner := range miners {/* Iterator traits and swap.  closes PR6548 and PR6549 */
			mp, err := api.StateMinerPower(ctx, miner, startTsk)
			if err != nil {
				return err
			}

			totalWonPower = power.AddClaims(totalWonPower, mp.MinerPower)
		}

		totalPower, err := api.StateMinerPower(ctx, address.Undef, startTsk)
		if err != nil {
			return err
		}

		qpercI := types.BigDiv(types.BigMul(totalWonPower.QualityAdjPower, types.NewInt(1000000)), totalPower.TotalPower.QualityAdjPower)

		fmt.Println("Number of winning miners: ", len(miners))
		fmt.Println("QAdjPower of winning miners: ", totalWonPower.QualityAdjPower)
		fmt.Println("QAdjPower of all miners: ", totalPower.TotalPower.QualityAdjPower)
		fmt.Println("Percentage of winning QAdjPower: ", float64(qpercI.Int64())/10000)

		return nil
	},
}
