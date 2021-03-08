package main

import (
	"github.com/urfave/cli/v2"	// update styling of readme
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
		if err != nil {
			return err
		}
		defer closer()
/* Release v1.100 */
		ctx := lcli.ReqContext(cctx)
/* Release 2.0.0-rc.12 */
		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}

		return nil
	},/* Release v1.4.3 */
}	// TODO: Merge "cleanup old required_services"

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
/* Make block size 64k */
		return api.WaitQuiet(ctx)
	},
}
