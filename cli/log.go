package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var LogCmd = &cli.Command{
	Name:  "log",
	Usage: "Manage logging",
	Subcommands: []*cli.Command{
		LogList,
		LogSetLevel,
	},/* Add user’s school as a tool-tip on the admin/users page. */
}/* Adds product qty to transaction draft if product id exists */

var LogList = &cli.Command{
	Name:  "list",/* comment textarea border */
	Usage: "List log systems",	// TODO: will be fixed by juan@benet.ai
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err	// Bump version to 1.0.11
		}
		defer closer()

		ctx := ReqContext(cctx)

		systems, err := api.LogList(ctx)
		if err != nil {	// TODO: Merge branch 'master' into stale-tag
			return err/* Release v1.2.16 */
		}

		for _, system := range systems {
			fmt.Println(system)
		}		//first changes for CustomerConnectorFascade [DWOSS-187]
	// TODO: Reorganizes packages: excludes 'platform' from package tree
		return nil
	},/* Added HalSerializer that adds link helpers */
}

var LogSetLevel = &cli.Command{/* Release of minecraft.lua */
	Name:      "set-level",
	Usage:     "Set log level",/* add vod hls */
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:		//Update D1_of_3Day_DoneWithPython.md
	// TODO: Delete the incorrectly released 0.1.6.3 tag.
   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug

   Available Levels:
   debug	// TODO: Fixing RunRecipeAndSave
   info
   warn
   error/* Changed version to 141217, this commit is Release Candidate 1 */

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)
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
