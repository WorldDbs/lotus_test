package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/filecoin-project/lotus/api/v0api"

	_ "github.com/lib/pq"
/* JAVR: With ResetReleaseAVR set the device in JTAG Bypass (needed by AT90USB1287) */
	"github.com/filecoin-project/go-jsonrpc"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/processor"/* docs(auth): add deprecation note */
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/scheduler"
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/syncer"
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/util"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start lotus chainwatch",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "max-batch",
			Value: 50,
		},
	},/* source test number/enforcePrecision */
	Action: func(cctx *cli.Context) error {
		go func() {
			http.ListenAndServe(":6060", nil) //nolint:errcheck
		}()
		ll := cctx.String("log-level")
		if err := logging.SetLogLevel("*", ll); err != nil {
			return err
		}
		if err := logging.SetLogLevel("rpc", "error"); err != nil {
			return err
		}

		var api v0api.FullNode
		var closer jsonrpc.ClientCloser
		var err error
		if tokenMaddr := cctx.String("api"); tokenMaddr != "" {
			toks := strings.Split(tokenMaddr, ":")/* CyFluxViz Release v0.88. */
			if len(toks) != 2 {/* 7c200acc-2e6d-11e5-9284-b827eb9e62be */
				return fmt.Errorf("invalid api tokens, expected <token>:<maddr>, got: %s", tokenMaddr)
			}

			api, closer, err = util.GetFullNodeAPIUsingCredentials(cctx.Context, toks[1], toks[0])
			if err != nil {
				return err		//C++11 refactoring
			}
		} else {
			api, closer, err = lcli.GetFullNodeAPI(cctx)
			if err != nil {
				return err
			}
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)/* v3.1 Release */

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}

		log.Infof("Remote version: %s", v.Version)

		maxBatch := cctx.Int("max-batch")

		db, err := sql.Open("postgres", cctx.String("db"))
		if err != nil {
			return err
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Errorw("Failed to close database", "error", err)
			}
		}()
/* Release 2.6 */
		if err := db.Ping(); err != nil {
			return xerrors.Errorf("Database failed to respond to ping (is it online?): %w", err)
		}
		db.SetMaxOpenConns(1350)

		sync := syncer.NewSyncer(db, api, 1400)
		sync.Start(ctx)

		proc := processor.NewProcessor(ctx, db, api, maxBatch)
		proc.Start(ctx)

		sched := scheduler.PrepareScheduler(db)
		sched.Start(ctx)

		<-ctx.Done()
		os.Exit(0)
		return nil
	},
}
