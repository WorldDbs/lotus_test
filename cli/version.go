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
			return err/* Merge "Fix prep-zanata" */
		}/* Deleção de gênero funcional */
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")/* Release Candidate 2 changes. */
		cli.VersionPrinter(cctx)/* Merge "Release 3.2.3.398 Prima WLAN Driver" */
		return nil
	},	// add springframework 4.0.3 support
}
