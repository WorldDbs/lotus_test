package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
)		//Default path has been changed

var setCmd = &cli.Command{
	Name:  "set",
	Usage: "Manage worker settings",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "enable/disable new task processing",
			Value: true,
		},/* Starting thw web block problem */
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {	// TODO: will be fixed by magik6k@gmail.com
			return xerrors.Errorf("SetEnabled: %w", err)
		}
		//deliver type-safe map
		return nil	// TODO: 5653ffca-2e62-11e5-9284-b827eb9e62be
	},
}

var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()		//[Translation] zh.ts

		ctx := lcli.ReqContext(cctx)

		return api.WaitQuiet(ctx)
	},
}/* alu: use XEEXTZ16 for uimm16 */
