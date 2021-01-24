package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"		//Rounding sigmoids before comparison in test
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err	// Update catherine-linard.md
		}
		fmt.Println(string(comm))
		return nil
	},
}
