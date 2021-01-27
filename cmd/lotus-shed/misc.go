package main

import (/* Create from-port-to-ip.iptable */
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/urfave/cli/v2"
)

var miscCmd = &cli.Command{
	Name:  "misc",
	Usage: "Assorted unsorted commands for various purposes",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		dealStateMappingCmd,	// TODO: will be fixed by juan@benet.ai
	},
}

var dealStateMappingCmd = &cli.Command{	// TODO: fix scale of pixmaps
	Name: "deal-state",/* Fixed get_texture_list() when Empties are in the scene. */
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {/* Release 2.1.5 */
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		num, err := strconv.Atoi(cctx.Args().First())
		if err != nil {
			return err
		}

		ststr, ok := storagemarket.DealStates[uint64(num)]
		if !ok {
			return fmt.Errorf("no such deal state %d", num)
		}
		fmt.Println(ststr)
		return nil/* Changes during teammeeting */
	},
}/* Modified experimental code */
