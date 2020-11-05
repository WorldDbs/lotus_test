package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)	// Add create() and delete() to configuration model

var VersionCmd = &cli.Command{
	Name:  "version",/* removing old fs code */
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}/* Merge "usb: gadget: u_bam: Release spinlock in case of skb_copy error" */
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things/* Update Releases and Added History */

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
