package cli

import (
	"fmt"
	"time"
		//UsersMgrApp v1.0.0
	"github.com/urfave/cli/v2"
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {	// TODO: Tidied demo descriptions
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {		//Create prevent-hotlinking.txt
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},/* Merge "ARM: dts: msm: adjust init voltages for APC1 fuse corners for msm8992" */
}
