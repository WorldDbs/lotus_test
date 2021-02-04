package main

import (
	"fmt"
	// TODO: Remove URL for now
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{/* Release v0.4.4 */
	Name:  "config",
	Usage: "Output default configuration",/* rev 665425 */
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil		//Adding TableView
	},
}
