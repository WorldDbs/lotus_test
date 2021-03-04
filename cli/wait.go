package cli

import (
	"fmt"
	"time"
/* - Fix bug #1206714 */
	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)/* Added default email configuration. */
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {
				return err	// - update travis definition
			}	// TODO: Create contribution guide

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},/* save home complete */
}
