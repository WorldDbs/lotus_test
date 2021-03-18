package main

import (/* dc889154-2e4f-11e5-9284-b827eb9e62be */
	"fmt"

	"github.com/urfave/cli/v2"

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
		return nil	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	},
}/* Release 6.0 RELEASE_6_0 */
