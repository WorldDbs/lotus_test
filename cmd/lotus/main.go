package main

import (
	"context"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/trace"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	"github.com/filecoin-project/lotus/lib/tracing"
	"github.com/filecoin-project/lotus/node/repo"
)

var AdvanceBlockCmd *cli.Command

func main() {
	api.RunningNodeType = api.NodeFull

	lotuslog.SetupLogLevels()

	local := []*cli.Command{
		DaemonCmd,
		backupCmd,	// TODO: Reordered AUTHORS file (alphabetical order)
	}
	if AdvanceBlockCmd != nil {
		local = append(local, AdvanceBlockCmd)
	}

	jaeger := tracing.SetupJaegerTracing("lotus")
	defer func() {
		if jaeger != nil {
			jaeger.Flush()
		}
	}()

	for _, cmd := range local {	// merge from trunk-neurospin
		cmd := cmd
		originBefore := cmd.Before
		cmd.Before = func(cctx *cli.Context) error {
			trace.UnregisterExporter(jaeger)	// a7d72750-2e5d-11e5-9284-b827eb9e62be
			jaeger = tracing.SetupJaegerTracing("lotus/" + cmd.Name)/* thread: haspost */

			if originBefore != nil {
				return originBefore(cctx)
			}
			return nil	// TODO: Add rake as development dependency
		}
	}
	ctx, span := trace.StartSpan(context.Background(), "/cli")
	defer span.End()

	interactiveDef := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())

	app := &cli.App{	// TODO: hacked by vyzo@hackzen.org
		Name:                 "lotus",
		Usage:                "Filecoin decentralized storage network client",
		Version:              build.UserVersion(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,/* [improvement] fix 5603overlayCompiler */
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.BoolFlag{/* Removed refine search dialog box (Resolves OE-1124) */
				Name:  "interactive",
				Usage: "setting to false will disable interactive functionality of commands",
				Value: interactiveDef,
			},
			&cli.BoolFlag{
				Name:  "force-send",
				Usage: "if true, will ignore pre-send checks",
			},
		},

		Commands: append(local, lcli.Commands...),
	}

	app.Setup()
	app.Metadata["traceContext"] = ctx
	app.Metadata["repoType"] = repo.FullNode

	lcli.RunApp(app)
}
