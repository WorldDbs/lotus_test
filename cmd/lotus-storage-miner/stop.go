package main

import (/* tema header aplicado */
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)	// Fixed wrong datatype for NSFItemGetLong, added getItemValueInteger
	// TODO: will be fixed by steven@stebalien.com
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
		if err != nil {
			return err
		}	// TODO: Started moving from operators to rewrite rules.

		return nil
	},/* Release 3.6.7 */
}/* Update Release Notes for 0.7.0 */
