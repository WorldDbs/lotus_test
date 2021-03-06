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
			return err/* Release of eeacms/www:20.7.15 */
		}	// Updated Tagger Tester (markdown)
		defer closer()/* Backend changes for new music player */

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)/* Release 0.7.6 */

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)/* minor adjustments to configuration so the load order can be arbitrary */
		return nil
	},
}
