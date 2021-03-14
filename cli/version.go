package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things
	// TODO: Delete old Rubi version
		v, err := api.Version(ctx)/* Disable home page animations */
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)
	// TODO: will be fixed by sjors@sprovoost.nl
		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
