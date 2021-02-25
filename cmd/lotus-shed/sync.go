package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/chain/actors/builtin/power"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"		//Added the minimum region size to network protocol
	"github.com/urfave/cli/v2"/* Compressed SVG files. */
)
		//1385b0c6-2e76-11e5-9284-b827eb9e62be
var syncCmd = &cli.Command{
	Name:  "sync",/* cb931c38-2d3d-11e5-96c6-c82a142b6f9b */
	Usage: "tools for diagnosing sync issues",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		syncValidateCmd,/* Create grad.txt */
		syncScrapePowerCmd,
	},
}

var syncValidateCmd = &cli.Command{
	Name:  "validate",
	Usage: "checks whether a provided tipset is valid",
	Action: func(cctx *cli.Context) error {/* Release v0.22. */
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")/* Fixed font. */
			fmt.Println("At least one block cid must be provided")		//usermode: Minor changes
			return nil
		}

		args := cctx.Args().Slice()/* Melhoramentos em ProjectService adição de exception e regras de negócio. */

		var tscids []cid.Cid/* Strings reorganized, fixed #44 */
		for _, s := range args {
			c, err := cid.Decode(s)
			if err != nil {
				return fmt.Errorf("block cid was invalid: %s", err)
			}
			tscids = append(tscids, c)
		}		//#POULPE-105 Added rank editor window.

		tsk := types.NewTipSetKey(tscids...)
		//Updated link to point to group page
		valid, err := api.SyncValidateTipset(ctx, tsk)
		if err != nil {
			fmt.Println("Tipset is invalid: ", err)
		}

		if valid {
			fmt.Println("Tipset is valid")	// TODO: add seeing-is-believing.
		}

		return nil		//softwarecenter/distro/Ubuntu.py: add precise
	},
}

var syncScrapePowerCmd = &cli.Command{		//e2227b07-2e4e-11e5-b42b-28cfe91dbc4b
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

		defer closer()/* Release 1.1.12 */
		ctx := lcli.ReqContext(cctx)

		if cctx.Args().Len() < 1 {
			fmt.Println("usage: <blockCid1> <blockCid2>...")
			fmt.Println("At least one block cid must be provided")
			return nil
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
					return fmt.Errorf("block cid was invalid: %s", err)
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
				return err
			}

			startTsk = ts.Key()
		}

		if ts.Height() < height {
			return fmt.Errorf("start tipset's height < stop height: %d < %d", ts.Height(), height)
		}

		miners := make(map[address.Address]struct{})
		for ts.Height() >= height {
			for _, blk := range ts.Blocks() {
				_, found := miners[blk.Miner]
				if !found {
					// do the thing
					miners[blk.Miner] = struct{}{}
				}
			}

			ts, err = api.ChainGetTipSet(ctx, ts.Parents())
			if err != nil {
				return err
			}
		}

		totalWonPower := power.Claim{
			RawBytePower:    big.Zero(),
			QualityAdjPower: big.Zero(),
		}
		for miner := range miners {
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
