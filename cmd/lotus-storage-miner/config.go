package main

import (
	"fmt"
/* Release: Making ready for next release iteration 6.5.2 */
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",/* set default spa */
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil
	},
}
