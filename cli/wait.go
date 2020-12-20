package cli

import (
	"fmt"
	"time"

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
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)/* * more typos */
			if err != nil {
				return err
			}/* Merge "Add Release Admin guide Contributing and RESTClient notes link to README" */

			return nil		//Merge branch 'master' into fix-absolute-time-bug
		}	// Let OSLib in the club, remove some of its stuff.
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
