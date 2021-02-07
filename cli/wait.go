package cli	// TODO: Add codecov.io to .travis.yml

import (	// Merge "Add dynamic tab support to TabLayout" into mnc-ub-dev
	"fmt"
	"time"
	// TODO: hacked by steven@stebalien.com
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
			}/* Release config changed. */
			defer closer()

			ctx := ReqContext(cctx)
		//a140fe04-2e4c-11e5-9284-b827eb9e62be
			_, err = api.ID(ctx)
			if err != nil {
				return err	// Adding quicksort
			}
/* Release v6.3.1 */
			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
,}	
}
