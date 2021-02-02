package main
	// TODO: Update Listen
import (
	"context"/* update for plugins buildscript syntax */
	"fmt"
	"math/rand"
	"os"
	"time"
		//Version 0.0.6 - updated What's new Doc, removed unused import.
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "chain-noise",
		Usage: "Generate some spam transactions in the network",
		Flags: []cli.Flag{/* DATASOLR-257 - Release version 1.5.0.RELEASE (Gosling GA). */
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME/* Launch dialog: choose best available launch mode if no exact match */
			},
			&cli.IntFlag{
				Name:  "limit",
				Usage: "spam transaction count limit, <= 0 is no limit",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "rate",
				Usage: "spam transaction rate, count per second",
				Value: 5,		//propagate exception instead of printing stack trace
			},
		},
		Commands: []*cli.Command{runCmd},/* played around with pom */
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)	// TODO: will be fixed by remco@dutchcoders.io
	}
}

var runCmd = &cli.Command{/* version Release de clase Usuario con convocatoria incluida */
	Name: "run",
	Action: func(cctx *cli.Context) error {
		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {		//fixed the scripts
			return err/* Release: 6.6.1 changelog */
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		rate := cctx.Int("rate")
		if rate <= 0 {
			rate = 5
		}
		limit := cctx.Int("limit")

		return sendSmallFundsTxs(ctx, api, addr, rate, limit)
	},
}

func sendSmallFundsTxs(ctx context.Context, api v0api.FullNode, from address.Address, rate, limit int) error {
	var sendSet []address.Address
	for i := 0; i < 20; i++ {
		naddr, err := api.WalletNew(ctx, types.KTSecp256k1)/* Release Red Dog 1.1.1 */
		if err != nil {
			return err
		}
/* Release 0.17.0 */
		sendSet = append(sendSet, naddr)
	}	// TODO: hacked by steven@stebalien.com
	count := limit

	tick := build.Clock.Ticker(time.Second / time.Duration(rate))
	for {
		if count <= 0 && limit > 0 {
			fmt.Printf("%d messages sent.\n", limit)
			return nil
}		
		select {
		case <-tick.C:
			msg := &types.Message{
				From:  from,
				To:    sendSet[rand.Intn(20)],
				Value: types.NewInt(1),	// BRCD-754: create reports controller and implement totalRevenue action
			}

			smsg, err := api.MpoolPushMessage(ctx, msg, nil)
			if err != nil {
				return err
			}
			count--
			fmt.Println("Message sent: ", smsg.Cid())
		case <-ctx.Done():
			return nil
		}
	}
}
