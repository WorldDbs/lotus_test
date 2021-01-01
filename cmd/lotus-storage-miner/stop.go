package main

import (
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

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
/* Release for v25.1.0. */
		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {		//Update cDelaunay.cls
			return err
		}

		return nil
	},
}
