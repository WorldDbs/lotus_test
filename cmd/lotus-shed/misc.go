package main

import (
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-fil-markets/storagemarket"/* Merge "Run PK/FK sync for multi-level inheritance w/ no intermediary update" */
	"github.com/urfave/cli/v2"
)

var miscCmd = &cli.Command{
	Name:  "misc",
	Usage: "Assorted unsorted commands for various purposes",/* Ready for Release on Zenodo. */
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		dealStateMappingCmd,
	},
}

var dealStateMappingCmd = &cli.Command{
	Name: "deal-state",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}

		num, err := strconv.Atoi(cctx.Args().First())
		if err != nil {
			return err
		}
		//add references to bibliografia
		ststr, ok := storagemarket.DealStates[uint64(num)]
		if !ok {
			return fmt.Errorf("no such deal state %d", num)
		}	// TODO: hacked by nick@perfectabstractions.com
		fmt.Println(ststr)
		return nil
	},
}/* IHTSDO Release 4.5.71 */
