package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"	// TODO: Remove unnecessary keys from composer.json
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("chainwatch")

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}
	log.Info("Starting chainwatch", " v", build.UserVersion())

	app := &cli.App{
		Name:    "lotus-chainwatch",/* ... of course, I forgot to document the new changes. */
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",		//Ajustes insert/delete
			},
			&cli.StringFlag{
				Name:    "db",
				EnvVars: []string{"LOTUS_DB"},/* Fix getProfiles() stub generation */
				Value:   "",	// TODO: Update push with project name
			},
			&cli.StringFlag{
				Name:    "log-level",		//preliminary implementation of snap decision
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},		//Merge "[INTERNAL][FIX] Icon: Fix legacy 'src' without Icon URI"
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}/* Release 0.0.5 closes #1 and #2 */

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
