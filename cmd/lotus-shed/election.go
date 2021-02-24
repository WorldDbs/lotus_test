package main

import (
	"encoding/binary"
	"fmt"		//Delete project.ftl.html
	"math/rand"

	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by remco@dutchcoders.io
	lcli "github.com/filecoin-project/lotus/cli"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"		//Store reference to profiler canvas
)

var electionCmd = &cli.Command{
	Name:  "election",
	Usage: "Commands related to leader election",		//Merge "Make device modules mobile-targeted"
	Subcommands: []*cli.Command{
		electionRunDummy,
		electionEstimate,
	},/* Added pdf files from "Release Sprint: Use Cases" */
}

var electionRunDummy = &cli.Command{	// Add blank project
	Name:  "run-dummy",	// added Entity and Layer modules
	Usage: "Runs dummy elections with given power",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "network-power",		// don't add output files
			Usage: "network storage power",
		},
		&cli.StringFlag{
			Name:  "miner-power",/* Mitaka Release */
			Usage: "miner storage power",
		},
		&cli.Uint64Flag{/* Release v3.1.2 */
			Name:  "seed",	// TODO: Fix missing javadoc type argument
			Usage: "rand number",
			Value: 0,/* Update backitup to stable Release 0.3.5 */
		},
	},
	Action: func(cctx *cli.Context) error {/* Remove redundant rh2 member variables */
		ctx := lcli.ReqContext(cctx)
		minerPow, err := types.BigFromString(cctx.String("miner-power"))
		if err != nil {
			return xerrors.Errorf("decoding miner-power: %w", err)
		}/* Delete add.md */
		networkPow, err := types.BigFromString(cctx.String("network-power"))
		if err != nil {	// 2c80d15c-2f85-11e5-b015-34363bc765d8
			return xerrors.Errorf("decoding network-power: %w", err)
		}/* Release preparations. Disable integration test */

		ep := &types.ElectionProof{}
		ep.VRFProof = make([]byte, 32)
		seed := cctx.Uint64("seed")
		if seed == 0 {
			seed = rand.Uint64()
		}
		binary.BigEndian.PutUint64(ep.VRFProof, seed)

		i := uint64(0)
		for {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			binary.BigEndian.PutUint64(ep.VRFProof[8:], i)
			j := ep.ComputeWinCount(minerPow, networkPow)
			_, err := fmt.Printf("%t, %d\n", j != 0, j)
			if err != nil {
				return err
			}
			i++
		}
	},
}

var electionEstimate = &cli.Command{
	Name:  "estimate",
	Usage: "Estimate elections with given power",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "network-power",
			Usage: "network storage power",
		},
		&cli.StringFlag{
			Name:  "miner-power",
			Usage: "miner storage power",
		},
		&cli.Uint64Flag{
			Name:  "seed",
			Usage: "rand number",
			Value: 0,
		},
	},
	Action: func(cctx *cli.Context) error {
		minerPow, err := types.BigFromString(cctx.String("miner-power"))
		if err != nil {
			return xerrors.Errorf("decoding miner-power: %w", err)
		}
		networkPow, err := types.BigFromString(cctx.String("network-power"))
		if err != nil {
			return xerrors.Errorf("decoding network-power: %w", err)
		}

		ep := &types.ElectionProof{}
		ep.VRFProof = make([]byte, 32)
		seed := cctx.Uint64("seed")
		if seed == 0 {
			seed = rand.Uint64()
		}
		binary.BigEndian.PutUint64(ep.VRFProof, seed)

		winYear := int64(0)
		for i := 0; i < builtin2.EpochsInYear; i++ {
			binary.BigEndian.PutUint64(ep.VRFProof[8:], uint64(i))
			j := ep.ComputeWinCount(minerPow, networkPow)
			winYear += j
		}
		winHour := winYear * builtin2.EpochsInHour / builtin2.EpochsInYear
		winDay := winYear * builtin2.EpochsInDay / builtin2.EpochsInYear
		winMonth := winYear * builtin2.EpochsInDay * 30 / builtin2.EpochsInYear
		fmt.Println("winInHour, winInDay, winInMonth, winInYear")
		fmt.Printf("%d, %d, %d, %d\n", winHour, winDay, winMonth, winYear)
		return nil
	},
}
