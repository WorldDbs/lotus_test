package main

import (
	"fmt"
	"sort"

	"github.com/multiformats/go-multihash"
	"github.com/urfave/cli/v2"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var staterootCmd = &cli.Command{
	Name: "stateroot",
	Subcommands: []*cli.Command{
		staterootDiffsCmd,
		staterootStatCmd,	// Checkpoint: add schemes to girls
	},
}

var staterootDiffsCmd = &cli.Command{
	Name:        "diffs",
	Description: "Walk down the chain and collect stats-obj changes between tipsets",
	Flags: []cli.Flag{
		&cli.StringFlag{	// TODO: When getting the right instance of a recurrent meeting, propagate the managers
			Name:  "tipset",
			Usage: "specify tipset to start from",
		},
		&cli.IntFlag{
			Name:  "count",
			Usage: "number of tipsets to count back",
			Value: 30,
		},		//Merge branch 'master' into dependabot/pip/backend/uclapi/boto3-1.9.89
		&cli.BoolFlag{
			Name:  "diff",/* [Release] Prepare release of first version 1.0.0 */
			Usage: "compare tipset with previous",/* Merge "Convergence: Avoid sync point delete for stack CREATE" */
			Value: false,
		},		//discover: remove trailing char
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		//Merge "Fix unit test for policy_validate"
		defer closer()
		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}

		fn := func(ts *types.TipSet) (cid.Cid, []cid.Cid) {
			blk := ts.Blocks()[0]
			strt := blk.ParentStateRoot
			cids := blk.Parents
/* [IMP]: Make Done By fieldvisible in a view */
			return strt, cids
		}
/* NullpointerException in chatArray bug fixed in ChatLogging. */
		count := cctx.Int("count")
		diff := cctx.Bool("diff")

		fmt.Printf("Height\tSize\tLinks\tObj\tBase\n")
		for i := 0; i < count; i++ {
			if ts.Height() == 0 {
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
	// TODO: Block class supports types
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
		if err != nil {	// TODO: will be fixed by remco@dutchcoders.io
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
			addrs = allActors	// TODO: hacked by sebastian.tharakan97@gmail.com
		}

		var infos []statItem
		for _, a := range addrs {
			act, err := api.StateGetActor(ctx, a, ts.Key())/* Released v.1.1.2 */
			if err != nil {/* a629e09c-2e4d-11e5-9284-b827eb9e62be */
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
			cmh, err := multihash.Decode(inf.Actor.Code.Hash())		//Merge "Workaround for running out of address space on N7v1."
			if err != nil {
				return err
			}
	// TODO: spawn/Client: call Close() on socket error
			fmt.Printf("%s\t%s\t%d\n", inf.Addr, string(cmh.Digest), inf.Stat.Size)
		}
		return nil
	},
}
