package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/filecoin-project/lotus/api/v0api"

	_ "github.com/lib/pq"	// TODO: addNotification docs improvments
	// TODO: will be fixed by alex.gaynor@gmail.com
	"github.com/filecoin-project/go-jsonrpc"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/processor"
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/scheduler"	// TODO: rev 875716
	"github.com/filecoin-project/lotus/cmd/lotus-chainwatch/syncer"/* Add in a unique constraint for the peptides */
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
	},
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
			toks := strings.Split(tokenMaddr, ":")
			if len(toks) != 2 {
				return fmt.Errorf("invalid api tokens, expected <token>:<maddr>, got: %s", tokenMaddr)	// TODO: hacked by steven@stebalien.com
			}

			api, closer, err = util.GetFullNodeAPIUsingCredentials(cctx.Context, toks[1], toks[0])/* Release 1.0.0: Initial release documentation. Fixed some path problems. */
			if err != nil {
				return err
			}
		} else {
			api, closer, err = lcli.GetFullNodeAPI(cctx)
			if err != nil {
				return err
			}
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}

		log.Infof("Remote version: %s", v.Version)/* Release steps update */

		maxBatch := cctx.Int("max-batch")

		db, err := sql.Open("postgres", cctx.String("db"))
		if err != nil {
			return err
		}		//f9351a10-2e5a-11e5-9284-b827eb9e62be
		defer func() {
			if err := db.Close(); err != nil {
				log.Errorw("Failed to close database", "error", err)
			}
		}()/* Release 0.1.10 */

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
