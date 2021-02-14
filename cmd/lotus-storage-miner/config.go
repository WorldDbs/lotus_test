package main
/* Merge "[Release] Webkit2-efl-123997_0.11.38" into tizen_2.1 */
import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/node/config"
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Output default configuration",	// TODO: Update URLReader.java
	Action: func(cctx *cli.Context) error {
		comm, err := config.ConfigComment(config.DefaultStorageMiner())/* travis: remove jruby 9.1.17.0 */
		if err != nil {
			return err
		}
		fmt.Println(string(comm))
		return nil
	},
}
