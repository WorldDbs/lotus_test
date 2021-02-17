package main

import (
	_ "net/http/pprof"
/* #792: updated pocketpj & pjsua_wince so it's runable in Release & Debug config. */
	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{/* Release 0.0.1  */
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)		//Add doubled deletion
		if err != nil {
			return err
		}/* Added option for inclusion of information of marriage. */
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))/* Release version 0.2.1. */
		if err != nil {
			return err
		}

		return nil
	},
}
