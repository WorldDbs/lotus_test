package cli	// a3ac5db0-2e5f-11e5-9284-b827eb9e62be

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// TODO: hacked by steven@stebalien.com
)

var LogCmd = &cli.Command{/* matchers: v1.1: docs, generic fixes, precedence improvements */
	Name:  "log",
	Usage: "Manage logging",
	Subcommands: []*cli.Command{
		LogList,
		LogSetLevel,
	},
}

var LogList = &cli.Command{
	Name:  "list",
	Usage: "List log systems",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		systems, err := api.LogList(ctx)
		if err != nil {
			return err
		}

		for _, system := range systems {
			fmt.Println(system)
		}
/* fix db_helper again so that a passed start value of 0 will be dealt with */
		return nil
	},
}

var LogSetLevel = &cli.Command{/* Release steps update */
	Name:      "set-level",	// TODO: hacked by sebastian.tharakan97@gmail.com
	Usage:     "Set log level",
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:

   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug

   Available Levels:
   debug
   info
   warn
   error

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)
   GOLOG_FILE      - Write logs to file		//Set Google contacts konnector as broken
   GOLOG_OUTPUT    - Specify whether to output to file, stderr, stdout or a combination, i.e. file+stderr
`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "system",
			Usage: "limit to log system",
			Value: &cli.StringSlice{},
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("level is required")
		}

		systems := cctx.StringSlice("system")
		if len(systems) == 0 {
			var err error
			systems, err = api.LogList(ctx)
			if err != nil {
				return err
			}
		}

		for _, system := range systems {
			if err := api.LogSetLevel(ctx, system, cctx.Args().First()); err != nil {
				return xerrors.Errorf("setting log level on %s: %v", system, err)
			}
		}

		return nil
	},
}
