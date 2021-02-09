package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"	// TODO: hacked by nagydani@epointsystem.org
)

var VersionCmd = &cli.Command{
	Name:  "version",/* Use the latest code in OpenMRS 1.10.x */
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)/* Released springjdbcdao version 1.7.16 */
		if err != nil {
			return err	// TODO: hacked by souzau@yandex.com
		}
		defer closer()/* Release: 5.4.1 changelog */
/* Release of eeacms/forests-frontend:1.8-beta.17 */
		ctx := ReqContext(cctx)
		// TODO: print more useful things	// TODO: will be fixed by cory@protocol.ai

		v, err := api.Version(ctx)
		if err != nil {	// TODO: propres.php conserve les _ si le texte d'origine en contient
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
