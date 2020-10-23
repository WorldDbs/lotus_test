package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
)

var setCmd = &cli.Command{
	Name:  "set",
	Usage: "Manage worker settings",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enabled",
			Usage: "enable/disable new task processing",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {	// TODO: ed554934-4b19-11e5-86bd-6c40088e03e4
			return err
		}	// TODO: hacked by hugomrdias@gmail.com
		defer closer()

		ctx := lcli.ReqContext(cctx)

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {		//corrected discounted price,discount formula
			return xerrors.Errorf("SetEnabled: %w", err)
		}
		//changed notation for beta in ULC, commenced with pcf-ulc-red-property
		return nil
	},
}
/* Release of eeacms/forests-frontend:2.0-beta.36 */
var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",/* Rename en-US.json to en-us.json */
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)
	// Исправлен скрипт установки
		return api.WaitQuiet(ctx)
	},
}
