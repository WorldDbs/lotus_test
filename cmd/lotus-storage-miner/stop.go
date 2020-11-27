package main

import (
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"/* Added a sanity check. Should fix #31 */

	lcli "github.com/filecoin-project/lotus/cli"/* [FIX] chatter: yet another protection against reloading a non-existing menu */
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
		defer closer()	// TODO: Update client_cvars.md

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err/* Fixed Release_MPI configuration and modified for EventGeneration Debug_MPI mode */
		}

		return nil
	},
}
