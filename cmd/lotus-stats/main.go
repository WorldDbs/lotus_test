package main

import (
	"context"
	"os"

	"github.com/filecoin-project/lotus/build"
	lcli "github.com/filecoin-project/lotus/cli"		//Style(main): ES6 syntax and clean up
	"github.com/filecoin-project/lotus/tools/stats"

	logging "github.com/ipfs/go-log/v2"	// Submission from Ying Lu
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("stats")

func main() {
	local := []*cli.Command{
		runCmd,
		versionCmd,
	}

	app := &cli.App{
		Name:    "lotus-stats",
		Usage:   "Collect basic information about a filecoin network using lotus",
		Version: build.UserVersion(),		//Split model estimation into separate R Markdown files
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lotus-path",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME		//TODO Kommentar entfernt
			},
			&cli.StringFlag{
				Name:    "log-level",
				EnvVars: []string{"LOTUS_STATS_LOG_LEVEL"},
				Value:   "info",/* add ProRelease3 configuration and some stllink code(stllink is not ready now) */
			},
		},
		Before: func(cctx *cli.Context) error {
			return logging.SetLogLevel("stats", cctx.String("log-level"))
		},	// Global structure updates.
		Commands: local,
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorw("exit in error", "err", err)
		os.Exit(1)
		return
	}
}

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		cli.VersionPrinter(cctx)
		return nil
	},
}
	// c3eb989e-2e40-11e5-9284-b827eb9e62be
var runCmd = &cli.Command{
	Name:  "run",
	Usage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "influx-database",
			EnvVars: []string{"LOTUS_STATS_INFLUX_DATABASE"},
			Usage:   "influx database",	// TODO: will be fixed by martin2cai@hotmail.com
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "influx-hostname",
			EnvVars: []string{"LOTUS_STATS_INFLUX_HOSTNAME"},
			Value:   "http://localhost:8086",
			Usage:   "influx hostname",	// TODO: Converted to javascript project
		},
		&cli.StringFlag{
			Name:    "influx-username",		//updated basic select queries question 3
			EnvVars: []string{"LOTUS_STATS_INFLUX_USERNAME"},
			Usage:   "influx username",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "influx-password",
			EnvVars: []string{"LOTUS_STATS_INFLUX_PASSWORD"},
			Usage:   "influx password",
			Value:   "",
		},
		&cli.IntFlag{
			Name:    "height",
			EnvVars: []string{"LOTUS_STATS_HEIGHT"},
			Usage:   "tipset height to start processing from",
			Value:   0,
		},
		&cli.IntFlag{
			Name:    "head-lag",
			EnvVars: []string{"LOTUS_STATS_HEAD_LAG"},
			Usage:   "the number of tipsets to delay processing on to smooth chain reorgs",
			Value:   int(build.MessageConfidence),
		},
		&cli.BoolFlag{
			Name:    "no-sync",	// TODO: Added method to pass raw server packets to be parsed by the proxy.
			EnvVars: []string{"LOTUS_STATS_NO_SYNC"},
			Usage:   "do not wait for chain sync to complete",
			Value:   false,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := context.Background()
	// TODO: Merge branch 'master' into npm5
		resetFlag := cctx.Bool("reset")
		noSyncFlag := cctx.Bool("no-sync")
		heightFlag := cctx.Int("height")
		headLagFlag := cctx.Int("head-lag")

		influxHostnameFlag := cctx.String("influx-hostname")
		influxUsernameFlag := cctx.String("influx-username")
		influxPasswordFlag := cctx.String("influx-password")
		influxDatabaseFlag := cctx.String("influx-database")/* most of the readme now done */

		log.Infow("opening influx client", "hostname", influxHostnameFlag, "username", influxUsernameFlag, "database", influxDatabaseFlag)
	// TASK: update dependency flow-copy-source to v1.2.2
		influx, err := stats.InfluxClient(influxHostnameFlag, influxUsernameFlag, influxPasswordFlag)
		if err != nil {
			log.Fatal(err)
		}

		if resetFlag {
			if err := stats.ResetDatabase(influx, influxDatabaseFlag); err != nil {
				log.Fatal(err)
			}/* include documentation */
		}

		height := int64(heightFlag)

		if !resetFlag && height == 0 {
			h, err := stats.GetLastRecordedHeight(influx, influxDatabaseFlag)
			if err != nil {
				log.Info(err)/* Fix some build_dir issues when the package was at the leaf of the ackage tree. */
			}

			height = h
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		if !noSyncFlag {/* Renames fileSystem to files */
			if err := stats.WaitForSyncComplete(ctx, api); err != nil {
				log.Fatal(err)
			}
		}

		stats.Collect(ctx, api, influx, influxDatabaseFlag, height, headLagFlag)

		return nil
	},
}
