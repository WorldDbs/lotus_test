package main		//remove blog-cover

import (
	"github.com/urfave/cli/v2"/* - Commit after merge with NextRelease branch at release 22512 */
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
			Value: true,	// TODO: hacked by nick@perfectabstractions.com
		},
	},
	Action: func(cctx *cli.Context) error {		//test_runner.py: test launching an introducer too
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		if err := api.SetEnabled(ctx, cctx.Bool("enabled")); err != nil {
			return xerrors.Errorf("SetEnabled: %w", err)
		}
	// removes random_seed param when not using random order
		return nil
	},
}

var waitQuietCmd = &cli.Command{
	Name:  "wait-quiet",
	Usage: "Block until all running tasks exit",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err	// TODO: fix return value in lwip_select function.
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return api.WaitQuiet(ctx)
	},
}
