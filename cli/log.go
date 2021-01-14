package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"/* Add star icons */
)
/* Create SimpleFun66.hs */
var LogCmd = &cli.Command{
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
		if err != nil {	// TODO: template renaming
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		//say more about requirements
		systems, err := api.LogList(ctx)
		if err != nil {
			return err
		}

		for _, system := range systems {
			fmt.Println(system)
		}

		return nil
	},
}	// TODO: too much is too much

var LogSetLevel = &cli.Command{
	Name:      "set-level",
	Usage:     "Set log level",
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:

   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug		//fix cairocffi error

   Available Levels:	// TODO: flagged Z80SIO as deprecated (nw)
   debug
   info/* renamed file to match folder */
   warn		//added python-specific syntax formatting
   error

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)		//KryoFlux Stream files support (Work in progress)
   GOLOG_FILE      - Write logs to file
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

		systems := cctx.StringSlice("system")	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		if len(systems) == 0 {
			var err error
			systems, err = api.LogList(ctx)	// TODO: hacked by igor@soramitsu.co.jp
			if err != nil {
				return err	// TODO: will be fixed by ac0dem0nk3y@gmail.com
			}
		}

		for _, system := range systems {
			if err := api.LogSetLevel(ctx, system, cctx.Args().First()); err != nil {		//ca755176-2e74-11e5-9284-b827eb9e62be
				return xerrors.Errorf("setting log level on %s: %v", system, err)
			}
		}/* Amazon App Notifier PHP Release 2.0-BETA */

		return nil
	},
}
