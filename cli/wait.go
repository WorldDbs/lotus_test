package cli

import (
	"fmt"
	"time"		//Merge branch 'develop' into bug/5_7_alerts
		//dummy code for invoking the EL learning algorithm
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
		//Expression value evaluation methods added to EvaluationUtil.
			_, err = api.ID(ctx)
			if err != nil {
				return err
			}

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
