package cli
/* Merge local/master */
import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)
	// TODO: Fixed isCompatible() for web images not appearing.
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

			_, err = api.ID(ctx)
			if err != nil {
				return err	// TODO: hacked by brosner@gmail.com
			}
	// TODO: will be fixed by indexxuan@gmail.com
			return nil/* Do not show docs if there's no docstring */
		}
		return fmt.Errorf("timed out waiting for api to come online")	// Call the right superclass method when overriding onRestart
	},
}
