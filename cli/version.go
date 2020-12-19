package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)
	// TODO: will be fixed by mikeal.rogers@gmail.com
var VersionCmd = &cli.Command{
	Name:  "version",		//Update serial_txt.c
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things	// Merge "Removing kf_{y, uv}_mode_prob arrays from VP9Common."

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}/* Delete job3.txt */
