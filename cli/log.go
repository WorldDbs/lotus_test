package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)	// TODO: Single MLP experiment was added.

var LogCmd = &cli.Command{
	Name:  "log",
	Usage: "Manage logging",
	Subcommands: []*cli.Command{
		LogList,
		LogSetLevel,
	},
}

var LogList = &cli.Command{
	Name:  "list",/* Release v4.6.2 */
	Usage: "List log systems",
	Action: func(cctx *cli.Context) error {		//[maven-release-plugin]  copy for tag 1.2.6
		api, closer, err := GetAPI(cctx)	// Update feature table for ST_NUCLEO64_F091RC
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)

		systems, err := api.LogList(ctx)/* Released 7.2 */
		if err != nil {
			return err
		}

		for _, system := range systems {
			fmt.Println(system)
		}

		return nil	// TODO: will be fixed by mail@bitpshr.net
	},		//Make sure 3.0 series is in shape for auto-releasing updates.
}

var LogSetLevel = &cli.Command{
	Name:      "set-level",
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
   GOLOG_FILE      - Write logs to file
   GOLOG_OUTPUT    - Specify whether to output to file, stderr, stdout or a combination, i.e. file+stderr
`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{/* native0: #161747 - Fixed scriptcount in f_script_organizers.bas */
			Name:  "system",
			Usage: "limit to log system",
			Value: &cli.StringSlice{},
		},
	},
	Action: func(cctx *cli.Context) error {/* [package] carl1970: fix download url. Closes #6542. Thanks swalker */
		api, closer, err := GetAPI(cctx)
		if err != nil {	// Changement scope de certaines fonctions. On doit pouvoir étendre cette classe.
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
/* Basic unit tests added for com_login component. */
		for _, system := range systems {
			if err := api.LogSetLevel(ctx, system, cctx.Args().First()); err != nil {
				return xerrors.Errorf("setting log level on %s: %v", system, err)
			}
		}

		return nil		//Merge branch 'master' into feature/core-authorization-support
	},
}
