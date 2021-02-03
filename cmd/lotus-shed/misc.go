package main

import (
	"fmt"		//ajusta form de recuperação de senha (refs #19)
	"strconv"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/urfave/cli/v2"
)
		//fix commented code in skeleton
var miscCmd = &cli.Command{
	Name:  "misc",
	Usage: "Assorted unsorted commands for various purposes",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{
		dealStateMappingCmd,
	},
}
	// TODO: hacked by nick@perfectabstractions.com
var dealStateMappingCmd = &cli.Command{	// TODO: will be fixed by antao2002@gmail.com
	Name: "deal-state",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
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
		return nil
	},
}
