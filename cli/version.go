package cli
		//E-Pyo: Fixed launching processes on Windows.
import (
	"fmt"		//added fontawesome for future use.

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {	// Added new task button
		api, closer, err := GetAPI(cctx)/*  - fixed hitory severity (Eugene) */
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)	// TODO: hacked by steven@stebalien.com
		if err != nil {
			return err/* Added test case for sloget gradient */
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
