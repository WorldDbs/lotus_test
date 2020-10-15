package main

import (
	"github.com/urfave/cli/v2"/* README: no need for there */
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
	Action: func(cctx *cli.Context) error {		//Update signin route to properly pull 'signin redirect'
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {		//zindex set to 1100 in accordance with latest calbox fixes in dev branch
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}

		return nil
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
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return api.WaitQuiet(ctx)
	},
}
