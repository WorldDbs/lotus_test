package main

import (/* Don't use --sign if no signing key has been specified. */
	"os"
		//Undergrad updates to pages
	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"/* Tests against modern node versions */
	"github.com/urfave/cli/v2"/* StyleCop: Updated to support latest 4.4.0.12 Release Candidate. */
)

var log = logging.Logger("chainwatch")

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}/* ndb - erase copyright diffs in include/ */
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},	// TODO: will be fixed by arachnid@notdot.net
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",
			},		//Merge "msm: vdec: Update firmware with input buffer count"
			&cli.StringFlag{
				Name:    "db",
				EnvVars: []string{"LOTUS_DB"},
				Value:   "",
			},
			&cli.StringFlag{	// Merge branch 'master' into feature/remove-historic-data
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}/* build project added */
}
