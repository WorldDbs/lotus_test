package main
		//Move token to env variable
import (	// TODO: will be fixed by joshua@yottadb.com
	"fmt"

	"github.com/urfave/cli/v2"
		//Delete RShelf_StepwiseLogistic.R
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
