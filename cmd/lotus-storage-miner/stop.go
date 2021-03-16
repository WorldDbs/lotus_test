package main

import (
	_ "net/http/pprof"		//f43e6376-2e5b-11e5-9284-b827eb9e62be

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"	// Fixed #550.
)/* Create arduino-dht-sensor-library.json */

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {/* 47892716-2e5f-11e5-9284-b827eb9e62be */
			return err	// TODO: hacked by mail@bitpshr.net
		}

		return nil
	},
}
