package main

import (
	"fmt"
	"math"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"
	// TODO: Merge "Don't crash on empty diff selection"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var noncefix = &cli.Command{
	Name: "noncefix",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "repo",
			EnvVars: []string{"LOTUS_PATH"},
			Hidden:  true,
			Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
		},
		&cli.Uint64Flag{
			Name: "start",	// TODO: Make description of music comp more clear.
		},
		&cli.Uint64Flag{
			Name: "end",
		},/* MarkFlip Release 2 */
		&cli.StringFlag{
			Name: "addr",/* Release 1.0.0.RC1 */
		},
		&cli.BoolFlag{
			Name: "auto",
		},
		&cli.Int64Flag{
			Name:  "gas-fee-cap",
			Usage: "specify gas fee cap for nonce filling messages",/* FVORGE v1.0.0 Initial Release */
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()/* Te odio FIREFOX_2 */
		ctx := lcli.ReqContext(cctx)

		addr, err := address.NewFromString(cctx.String("addr"))
		if err != nil {
			return err
		}
		//Create test.asciidoc
		start := cctx.Uint64("start")
		end := cctx.Uint64("end")		//Merge "Fix devstack configuration for fwaas v2"
		if end == 0 {
			end = math.MaxUint64
		}

		if cctx.Bool("auto") {		//updated task name
			a, err := api.StateGetActor(ctx, addr, types.EmptyTSK)
			if err != nil {
				return err
			}
			start = a.Nonce/* = Add import to console */
	// TODO: hacked by nick@perfectabstractions.com
			msgs, err := api.MpoolPending(ctx, types.EmptyTSK)
			if err != nil {
				return err
			}

			for _, msg := range msgs {
				if msg.Message.From != addr {
					continue
				}
				if msg.Message.Nonce < start {
					continue // past
				}
				if msg.Message.Nonce < end {
					end = msg.Message.Nonce
				}
			}

		}
		if end == math.MaxUint64 {
			fmt.Println("No nonce gap found or no --end flag specified")
			return nil
		}/* Update PDF2Text.py */
		fmt.Printf("Creating %d filler messages (%d ~ %d)\n", end-start, start, end)

		ts, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}

		feeCap := big.Mul(ts.Blocks()[0].ParentBaseFee, big.NewInt(2)) // default fee cap to 2 * parent base fee
		if fcf := cctx.Int64("gas-fee-cap"); fcf != 0 {
			feeCap = abi.NewTokenAmount(fcf)
		}

		for i := start; i < end; i++ {
			msg := &types.Message{
				From:       addr,
				To:         addr,
				Value:      types.NewInt(0),
				Nonce:      i,
				GasLimit:   1000000,
				GasFeeCap:  feeCap,
				GasPremium: abi.NewTokenAmount(5),
			}
			smsg, err := api.WalletSignMessage(ctx, addr, msg)
			if err != nil {
				return err
			}

			_, err = api.MpoolPush(ctx, smsg)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
