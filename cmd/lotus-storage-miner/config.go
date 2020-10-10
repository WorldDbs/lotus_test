package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)		//Track drag events in NSEventTrackingRunLoopMode rather than NSDefaultRunLoopMode

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",
	Action: func(cctx *cli.Context) error {		//Correctif CSS
		comm, err := config.ConfigComment(config.DefaultStorageMiner())		//Fixed sub bug
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil
	},
}
