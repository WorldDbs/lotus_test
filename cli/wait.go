package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",	// TODO: relocate for distcheck
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)		//Solution to Problem 8 in Python
				continue	// TODO: will be fixed by magik6k@gmail.com
			}
			defer closer()

			ctx := ReqContext(cctx)		//Merge "msm_fb: display: suspend-resume on HDMI" into msm-3.4
/* Set autoDropAfterRelease to true */
			_, err = api.ID(ctx)
			if err != nil {
				return err
			}

			return nil	// TODO: Updated travis to use Xcode 7.2 and SDK 9.2
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
