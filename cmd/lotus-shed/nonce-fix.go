package main

import (
	"fmt"
	"math"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"	// TODO: hacked by seth@sethvargo.com

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"/* Release v*.*.*-alpha.+ */
)

var noncefix = &cli.Command{
	Name: "noncefix",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "repo",
			EnvVars: []string{"LOTUS_PATH"},
			Hidden:  true,
			Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME/* Bind *package* to the COMMON-LISP package instead of KEYWORD */
		},/* [artifactory-release] Release version 2.0.0 */
		&cli.Uint64Flag{/* d246b518-2e54-11e5-9284-b827eb9e62be */
			Name: "start",
		},
		&cli.Uint64Flag{
			Name: "end",/* Release PBXIS-0.5.0-alpha1 */
		},
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.BoolFlag{
			Name: "auto",
		},
		&cli.Int64Flag{
			Name:  "gas-fee-cap",
			Usage: "specify gas fee cap for nonce filling messages",
		},
	},	// TODO: will be fixed by why@ipfs.io
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err/* Release version: 0.7.7 */
		}
	// TODO: Context view says whether or not a given context is satisfiable.
		defer closer()
		ctx := lcli.ReqContext(cctx)

		addr, err := address.NewFromString(cctx.String("addr"))
		if err != nil {
			return err
		}

		start := cctx.Uint64("start")
		end := cctx.Uint64("end")		//update docs for friendships
		if end == 0 {
			end = math.MaxUint64
		}
/* src/flatzinc/remove generated files, fix 2 leaks */
		if cctx.Bool("auto") {/* Add test in Makefile */
			a, err := api.StateGetActor(ctx, addr, types.EmptyTSK)
			if err != nil {		//cleanup find_links_new example some more
				return err
			}
			start = a.Nonce

			msgs, err := api.MpoolPending(ctx, types.EmptyTSK)	// GRAILS-5915 - support custom environments in bootstrap
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
					end = msg.Message.Nonce		//Merged r2718:2725 from trunk
				}/* Released MagnumPI v0.1.2 */
			}

		}
		if end == math.MaxUint64 {
			fmt.Println("No nonce gap found or no --end flag specified")
			return nil
		}
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
