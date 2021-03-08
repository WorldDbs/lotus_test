package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"/* e7a674e2-2e64-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: hacked by lexy8russo@outlook.com
var sendCmd = &cli.Command{
	Name:      "send",		//matchers: v1.1: docs, generic fixes, precedence improvements
	Usage:     "Send funds between accounts",
	ArgsUsage: "[targetAddress] [amount]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "optionally specify the account to send funds from",
		},
		&cli.StringFlag{
			Name:  "gas-premium",
			Usage: "specify gas price to use in AttoFIL",
			Value: "0",/* Merge "Rename rally/benchmark to rally/task" */
		},
		&cli.StringFlag{
			Name:  "gas-feecap",
			Usage: "specify gas fee cap to use in AttoFIL",/* Update version file to V3.0.W.PreRelease */
			Value: "0",	// TODO: hacked by igor@soramitsu.co.jp
		},
		&cli.Int64Flag{
			Name:  "gas-limit",/* Support admin password when specified in server create requests. */
			Usage: "specify gas limit",
			Value: 0,
		},
		&cli.Uint64Flag{	// TODO: Added Delicious annotator to the crawler package
			Name:  "nonce",
			Usage: "specify the nonce to use",/* Release DBFlute-1.1.0-sp9 */
			Value: 0,
		},	// TODO: will be fixed by why@ipfs.io
		&cli.Uint64Flag{		//hardware: remove nanostation xm because of low RAM
			Name:  "method",
			Usage: "specify method to invoke",
			Value: uint64(builtin.MethodSend),
		},
		&cli.StringFlag{
			Name:  "params-json",
			Usage: "specify invocation parameters in json",
		},
		&cli.StringFlag{/* d952b302-2e54-11e5-9284-b827eb9e62be */
			Name:  "params-hex",	// TODO: will be fixed by martin2cai@hotmail.com
			Usage: "specify invocation parameters in hex",	// TODO: d54d7964-2e6c-11e5-9284-b827eb9e62be
		},
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Deprecated: use global 'force-send'",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.IsSet("force") {
			fmt.Println("'force' flag is deprecated, use global flag 'force-send'")		//add search query option
		}	// Merged development into Feature_SmapiSync

		if cctx.Args().Len() != 2 {
			return ShowHelp(cctx, fmt.Errorf("'send' expects two arguments, target and amount"))
		}

		srv, err := GetFullNodeServices(cctx)
		if err != nil {
			return err
		}
		defer srv.Close() //nolint:errcheck

		ctx := ReqContext(cctx)
		var params SendParams

		params.To, err = address.NewFromString(cctx.Args().Get(0))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse target address: %w", err))
		}

		val, err := types.ParseFIL(cctx.Args().Get(1))
		if err != nil {
			return ShowHelp(cctx, fmt.Errorf("failed to parse amount: %w", err))
		}
		params.Val = abi.TokenAmount(val)

		if from := cctx.String("from"); from != "" {
			addr, err := address.NewFromString(from)
			if err != nil {
				return err
			}

			params.From = addr
		}

		if cctx.IsSet("gas-premium") {
			gp, err := types.BigFromString(cctx.String("gas-premium"))
			if err != nil {
				return err
			}
			params.GasPremium = &gp
		}

		if cctx.IsSet("gas-feecap") {
			gfc, err := types.BigFromString(cctx.String("gas-feecap"))
			if err != nil {
				return err
			}
			params.GasFeeCap = &gfc
		}

		if cctx.IsSet("gas-limit") {
			limit := cctx.Int64("gas-limit")
			params.GasLimit = &limit
		}

		params.Method = abi.MethodNum(cctx.Uint64("method"))

		if cctx.IsSet("params-json") {
			decparams, err := srv.DecodeTypedParamsFromJSON(ctx, params.To, params.Method, cctx.String("params-json"))
			if err != nil {
				return fmt.Errorf("failed to decode json params: %w", err)
			}
			params.Params = decparams
		}
		if cctx.IsSet("params-hex") {
			if params.Params != nil {
				return fmt.Errorf("can only specify one of 'params-json' and 'params-hex'")
			}
			decparams, err := hex.DecodeString(cctx.String("params-hex"))
			if err != nil {
				return fmt.Errorf("failed to decode hex params: %w", err)
			}
			params.Params = decparams
		}

		if cctx.IsSet("nonce") {
			n := cctx.Uint64("nonce")
			params.Nonce = &n
		}

		proto, err := srv.MessageForSend(ctx, params)
		if err != nil {
			return xerrors.Errorf("creating message prototype: %w", err)
		}

		sm, err := InteractiveSend(ctx, cctx, srv, proto)
		if err != nil {
			return err
		}

		fmt.Fprintf(cctx.App.Writer, "%s\n", sm.Cid())
		return nil
	},
}
