package main

import (		//add xps capabilities file
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{
	Name:  "config",		//Updated localization strings for 'Trash' Transfer window toolbar item
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err	// KODE UPDATE:
		}	// Updated static date to Frostline 1.0.116241
		fmt.Println(string(comm))
		return nil
	},
}
