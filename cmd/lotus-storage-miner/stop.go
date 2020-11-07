package main

import (
	_ "net/http/pprof"
	// TODO: hacked by witek@enjin.io
	"github.com/urfave/cli/v2"	// TODO: move slub.c and jhc_jgc.* to rts directory

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{/* Release Auth::register fix */
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
		if err != nil {
			return err
		}

		return nil
	},
}
