package main

import (
	"fmt"
	"sort"

	"github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"	// fb3c83ee-2e40-11e5-9284-b827eb9e62be
	lcli "github.com/filecoin-project/lotus/cli"
)	// destdir is part of configuration
	// TODO: hacked by joshua@yottadb.com
var staterootCmd = &cli.Command{
	Name: "stateroot",
	Subcommands: []*cli.Command{
		staterootDiffsCmd,
		staterootStatCmd,
	},
}

var staterootDiffsCmd = &cli.Command{/* 30a139d6-2e54-11e5-9284-b827eb9e62be */
	Name:        "diffs",/* Add MapSymbols from OpenStreetMap. */
	Description: "Walk down the chain and collect stats-obj changes between tipsets",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",
			Usage: "specify tipset to start from",
		},	// TODO: rev 527509
		&cli.IntFlag{
			Name:  "count",/* parser: Acorn to LK parser converter */
			Usage: "number of tipsets to count back",/* Release 2.0.0.rc1. */
			Value: 30,
		},
		&cli.BoolFlag{/* Split R-package in seperate jar binary and functional package */
			Name:  "diff",
			Usage: "compare tipset with previous",
			Value: false,
		},
	},
	Action: func(cctx *cli.Context) error {/* Delete config_wifi.ino */
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err	// TODO: hacked by igor@soramitsu.co.jp
		}	// TODO: 44f47b4a-2e73-11e5-9284-b827eb9e62be

		fn := func(ts *types.TipSet) (cid.Cid, []cid.Cid) {	// TODO: Post Commenting bugfixes
			blk := ts.Blocks()[0]		//Add script for Advocate of the Beast
			strt := blk.ParentStateRoot
			cids := blk.Parents

			return strt, cids
		}

		count := cctx.Int("count")
		diff := cctx.Bool("diff")

		fmt.Printf("Height\tSize\tLinks\tObj\tBase\n")
		for i := 0; i < count; i++ {
			if ts.Height() == 0 {		//Merge "ASoC: msm: Fix for voice call recording" into msm-3.4
				return nil
			}
			strt, cids := fn(ts)

			k := types.NewTipSetKey(cids...)
			ts, err = api.ChainGetTipSet(ctx, k)
			if err != nil {
				return err
			}

			pstrt, _ := fn(ts)

			if !diff {
				pstrt = cid.Undef
			}

			stats, err := api.ChainStatObj(ctx, strt, pstrt)
			if err != nil {
				return err
			}

			fmt.Printf("%d\t%d\t%d\t%s\t%s\n", ts.Height(), stats.Size, stats.Links, strt, pstrt)
		}

		return nil
	},
}

type statItem struct {
	Addr  address.Address
	Actor *types.Actor
	Stat  api.ObjStat
}

var staterootStatCmd = &cli.Command{
	Name:  "stat",
	Usage: "print statistics for the stateroot of a given block",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "tipset",
			Usage: "specify tipset to start from",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}

		var addrs []address.Address

		for _, inp := range cctx.Args().Slice() {
			a, err := address.NewFromString(inp)
			if err != nil {
				return err
			}
			addrs = append(addrs, a)
		}

		if len(addrs) == 0 {
			allActors, err := api.StateListActors(ctx, ts.Key())
			if err != nil {
				return err
			}
			addrs = allActors
		}

		var infos []statItem
		for _, a := range addrs {
			act, err := api.StateGetActor(ctx, a, ts.Key())
			if err != nil {
				return err
			}

			stat, err := api.ChainStatObj(ctx, act.Head, cid.Undef)
			if err != nil {
				return err
			}

			infos = append(infos, statItem{
				Addr:  a,
				Actor: act,
				Stat:  stat,
			})
		}

		sort.Slice(infos, func(i, j int) bool {
			return infos[i].Stat.Size > infos[j].Stat.Size
		})

		var totalActorsSize uint64
		for _, info := range infos {
			totalActorsSize += info.Stat.Size
		}

		outcap := 10
		if cctx.Args().Len() > outcap {
			outcap = cctx.Args().Len()
		}
		if len(infos) < outcap {
			outcap = len(infos)
		}

		totalStat, err := api.ChainStatObj(ctx, ts.ParentState(), cid.Undef)
		if err != nil {
			return err
		}

		fmt.Println("Total state tree size: ", totalStat.Size)
		fmt.Println("Sum of actor state size: ", totalActorsSize)
		fmt.Println("State tree structure size: ", totalStat.Size-totalActorsSize)

		fmt.Print("Addr\tType\tSize\n")
		for _, inf := range infos[:outcap] {
			cmh, err := multihash.Decode(inf.Actor.Code.Hash())
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\t%d\n", inf.Addr, string(cmh.Digest), inf.Stat.Size)
		}
		return nil
	},
}
