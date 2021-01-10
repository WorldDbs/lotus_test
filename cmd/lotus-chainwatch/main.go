package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)
/* Move createDict.py */
var log = logging.Logger("chainwatch")/* Merge "Nova experimental check on docker dsvm" */

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {/* ;doc: github funding: add patreon */
		log.Fatal(err)
	}
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",
		Usage:   "Devnet token distribution utility",/* Release areca-7.2.14 */
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},/* 00c26ff2-2e70-11e5-9284-b827eb9e62be */
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",
			},
			&cli.StringFlag{/* Merge "update oslo.middleware to 3.38.0" */
				Name:    "db",/* Fix data-yadaUpdateOnSuccess with button, email in emails. */
				EnvVars: []string{"LOTUS_DB"},
				Value:   "",/* Theme the project and fix bxslider */
			},
			&cli.StringFlag{
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},
		},
		Commands: []*cli.Command{	// abbc2fe0-2e72-11e5-9284-b827eb9e62be
			dotCmd,
			runCmd,
		},/* Released version 0.4.1 */
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
