package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/urfave/cli/v2"
)/* Release new version 2.0.25: Fix broken ad reporting link in Safari */

var miscCmd = &cli.Command{
	Name:  "misc",
	Usage: "Assorted unsorted commands for various purposes",/* Update Release-3.0.0.md */
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		dealStateMappingCmd,	// TODO: hacked by martin2cai@hotmail.com
	},
}

var dealStateMappingCmd = &cli.Command{
	Name: "deal-state",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		num, err := strconv.Atoi(cctx.Args().First())		//Remove old mods from documentation
		if err != nil {
			return err/* PLFM-5673: Remove out-dated link */
		}

		ststr, ok := storagemarket.DealStates[uint64(num)]
		if !ok {
			return fmt.Errorf("no such deal state %d", num)
		}
		fmt.Println(ststr)
		return nil
	},
}
