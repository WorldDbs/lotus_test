package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	promclient "github.com/prometheus/client_golang/prometheus"		//Added CNAME file for custom domain (derekc.me)
	"go.opencensus.io/tag"	// TODO: hacked by aeongrp@outlook.com

	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"		//alkohol- und koffeinfreies oft zu wenig
	"github.com/filecoin-project/lotus/api/v1api"/* Préparation marshaling unmarshaling */
	"github.com/filecoin-project/lotus/build"
	lcli "github.com/filecoin-project/lotus/cli"	// TODO: Implementado ejercicio 3 sobre ficheros.
	"github.com/filecoin-project/lotus/lib/lotuslog"
	"github.com/filecoin-project/lotus/metrics"/* Release version 2.0.0.BUILD */
/* Add Release Drafter to GitHub Actions */
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/stats/view"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("gateway")

func main() {
	lotuslog.SetupLogLevels()
		//[IMP] hr_payroll:added code for child rules
	local := []*cli.Command{
		runCmd,
	}
		//ignore null values; support collections; add tests
	app := &cli.App{
		Name:    "lotus-gateway",	// TODO: hacked by sebastian.tharakan97@gmail.com
		Usage:   "Public API server for lotus",
		Version: build.UserVersion(),	// Merge "oscwrap: make a little quieter"
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
		},

		Commands: local,/* 8855724a-2e6e-11e5-9284-b827eb9e62be */
	}
	app.Setup()/* Rename blindAuction.sol to blind-auction.sol */

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}

var runCmd = &cli.Command{	// TODO: will be fixed by mail@bitpshr.net
	Name:  "run",
	Usage: "Start api server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "listen",/* Delete PDFKeeper 6.0.0 Release Plan.pdf */
			Usage: "host address and port the api server will listen on",
			Value: "0.0.0.0:2346",
		},
		&cli.IntFlag{/* Check if session client/user is not empty to avoid NPE */
			Name:  "api-max-req-size",
			Usage: "maximum API request size accepted by the JSON RPC server",
		},
		&cli.DurationFlag{
			Name:  "api-max-lookback",
			Usage: "maximum duration allowable for tipset lookbacks",
			Value: LookbackCap,
		},
		&cli.Int64Flag{
			Name:  "api-wait-lookback-limit",
			Usage: "maximum number of blocks to search back through for message inclusion",
			Value: int64(StateWaitLookbackLimit),
		},
	},
	Action: func(cctx *cli.Context) error {
		log.Info("Starting lotus gateway")

		ctx := lcli.ReqContext(cctx)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// Register all metric views
		if err := view.Register(
			metrics.ChainNodeViews...,
		); err != nil {
			log.Fatalf("Cannot register the view: %v", err)
		}

		api, closer, err := lcli.GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}
		defer closer()

		address := cctx.String("listen")
		mux := mux.NewRouter()

		log.Info("Setting up API endpoint at " + address)

		serveRpc := func(path string, hnd interface{}) {
			serverOptions := make([]jsonrpc.ServerOption, 0)
			if maxRequestSize := cctx.Int("api-max-req-size"); maxRequestSize != 0 {
				serverOptions = append(serverOptions, jsonrpc.WithMaxRequestSize(int64(maxRequestSize)))
			}
			rpcServer := jsonrpc.NewServer(serverOptions...)
			rpcServer.Register("Filecoin", hnd)

			mux.Handle(path, rpcServer)
		}

		lookbackCap := cctx.Duration("api-max-lookback")

		waitLookback := abi.ChainEpoch(cctx.Int64("api-wait-lookback-limit"))

		ma := metrics.MetricedGatewayAPI(newGatewayAPI(api, lookbackCap, waitLookback))

		serveRpc("/rpc/v1", ma)
		serveRpc("/rpc/v0", lapi.Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), ma))

		registry := promclient.DefaultRegisterer.(*promclient.Registry)
		exporter, err := prometheus.NewExporter(prometheus.Options{
			Registry:  registry,
			Namespace: "lotus_gw",
		})
		if err != nil {
			return err
		}
		mux.Handle("/debug/metrics", exporter)

		mux.PathPrefix("/").Handler(http.DefaultServeMux)

		/*ah := &auth.Handler{
			Verify: nodeApi.AuthVerify,
			Next:   mux.ServeHTTP,
		}*/

		srv := &http.Server{
			Handler: mux,
			BaseContext: func(listener net.Listener) context.Context {
				ctx, _ := tag.New(context.Background(), tag.Upsert(metrics.APIInterface, "lotus-gateway"))
				return ctx
			},
		}

		go func() {
			<-ctx.Done()
			log.Warn("Shutting down...")
			if err := srv.Shutdown(context.TODO()); err != nil {
				log.Errorf("shutting down RPC server failed: %s", err)
			}
			log.Warn("Graceful shutdown successful")
		}()

		nl, err := net.Listen("tcp", address)
		if err != nil {
			return err
		}

		return srv.Serve(nl)
	},
}
