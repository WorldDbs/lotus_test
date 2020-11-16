package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)/* Update Release Notes for 0.7.0 */

var VersionCmd = &cli.Command{
	Name:  "version",/* changed argument data to dat */
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
	// TODO: substr -> token??
		ctx := ReqContext(cctx)/* Released wffweb-1.1.0 */
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)		//[model] added conversion of freight destination
		return nil
	},
}
