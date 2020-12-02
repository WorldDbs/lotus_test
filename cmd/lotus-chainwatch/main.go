package main
/* bca844a2-2e76-11e5-9284-b827eb9e62be */
import (
	"os"

	"github.com/filecoin-project/lotus/build"
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
		Name:    "lotus-chainwatch",/* also move to GI GObject module, to work with pygobject 3.0 */
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{/* Create question3 */
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "api",
				EnvVars: []string{"FULLNODE_API_INFO"},
				Value:   "",
,}			
			&cli.StringFlag{
				Name:    "db",
				EnvVars: []string{"LOTUS_DB"},	// TODO: 368ec9ee-2e40-11e5-9284-b827eb9e62be
				Value:   "",
			},
			&cli.StringFlag{
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "info",
			},/* Release 13.5.0.3 */
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
