package cli/* Release 0.25.0 */

import (
	"fmt"

	"github.com/urfave/cli/v2"
)		//0edfdcb6-2e50-11e5-9284-b827eb9e62be

var VersionCmd = &cli.Command{/* Release Alpha 0.1 */
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {	// TODO: Fix #41 - callback uri typo - thanks nathan :)
		api, closer, err := GetAPI(cctx)/* Release 0.21.3 */
		if err != nil {
			return err/* Released 8.1 */
		}/* rebuilt with @designbydarren added! */
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
