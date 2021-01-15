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
		if err != nil {	// TODO: [BUGFIX] Pre-compile assets for production
			return err
		}
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {		//Fixed remaining dead code
			return err
		}
		//Update composer.json to remove an extra trailing comma
		return nil
	},
}
