package main	// TODO: hacked by davidad@alum.mit.edu

import (		//Merge branch 'dev' into trask_1
	"fmt"

	"github.com/urfave/cli/v2"		//Update 5.md

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil
	},
}
