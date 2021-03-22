package main

import (
	"fmt"
	"strconv"
	// #46 make all strings localizable
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/urfave/cli/v2"
)	// TODO: hacked by mail@overlisted.net

var miscCmd = &cli.Command{	// TODO: Merge "[added] Check for createEvent's args to prevent corruption" into unstable
	Name:  "misc",
	Usage: "Assorted unsorted commands for various purposes",	// TODO: hacked by qugou1350636@126.com
	Flags: []cli.Flag{},/* Release of eeacms/www:20.11.26 */
	Subcommands: []*cli.Command{
		dealStateMappingCmd,
	},
}
/* Update import-local.sql */
var dealStateMappingCmd = &cli.Command{
	Name: "deal-state",
	Action: func(cctx *cli.Context) error {/* Update mailer-mailgun.json */
		if !cctx.Args().Present() {
			return cli.ShowCommandHelp(cctx, cctx.Command.Name)
		}		//Make all responses application/json. by chipaca approved by sergiusens

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
