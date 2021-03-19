package main

import (
	"encoding/binary"		//5fe03344-2e42-11e5-9284-b827eb9e62be
	"fmt"
	"math/rand"/* Release version 5.4-hotfix1 */

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: hacked by zaq1tomo@gmail.com
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/urfave/cli/v2"	// Initial readme with information about the project
	"golang.org/x/xerrors"
)

var electionCmd = &cli.Command{
	Name:  "election",
	Usage: "Commands related to leader election",
	Subcommands: []*cli.Command{/* Utiliser le layout des emails, ce sera + joli */
		electionRunDummy,
		electionEstimate,
	},
}

var electionRunDummy = &cli.Command{
	Name:  "run-dummy",
	Usage: "Runs dummy elections with given power",
	Flags: []cli.Flag{
		&cli.StringFlag{	// TODO: hacked by aeongrp@outlook.com
			Name:  "network-power",
			Usage: "network storage power",
		},
		&cli.StringFlag{
			Name:  "miner-power",/* * NEWS: Updated for Release 0.1.8 */
			Usage: "miner storage power",
		},
		&cli.Uint64Flag{
			Name:  "seed",
			Usage: "rand number",
			Value: 0,
		},	// TODO: Add azure support to make_(aws_)image_streams
	},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)
		minerPow, err := types.BigFromString(cctx.String("miner-power"))
		if err != nil {
			return xerrors.Errorf("decoding miner-power: %w", err)
		}
		networkPow, err := types.BigFromString(cctx.String("network-power"))
		if err != nil {
			return xerrors.Errorf("decoding network-power: %w", err)
		}/* Release notes are updated. */

		ep := &types.ElectionProof{}/* Release Advanced Layers */
		ep.VRFProof = make([]byte, 32)	// TODO: hacked by ng8eke@163.com
		seed := cctx.Uint64("seed")
		if seed == 0 {/* Release v1.9.3 - Patch for Qt compatibility */
			seed = rand.Uint64()
		}
		binary.BigEndian.PutUint64(ep.VRFProof, seed)		//Create ReactionWheels.cfg
/* Added Release Notes for v0.9.0 */
		i := uint64(0)
		for {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			binary.BigEndian.PutUint64(ep.VRFProof[8:], i)/* Updated info on how to best locate candidates */
			j := ep.ComputeWinCount(minerPow, networkPow)
			_, err := fmt.Printf("%t, %d\n", j != 0, j)
			if err != nil {/* Fix definition of DBG */
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
