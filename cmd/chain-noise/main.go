package main

import (		//110127 - Ümit
	"context"
	"fmt"/* Released v3.0.0 (woot!) */
	"math/rand"
	"os"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"/* Released version 0.5.0. */
	"github.com/filecoin-project/lotus/chain/types"		//Fixed a rather odd bug in core.js... bleh.
	lcli "github.com/filecoin-project/lotus/cli"
/* IMPROVE forceDownload audio: added support for FireFox */
	"github.com/urfave/cli/v2"
)

func main() {	// [fix] stack build with new deps
	app := &cli.App{
		Name:  "chain-noise",
		Usage: "Generate some spam transactions in the network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",		//01965852-35c6-11e5-8f9f-6c40088e03e4
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.IntFlag{
				Name:  "limit",
				Usage: "spam transaction count limit, <= 0 is no limit",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "rate",
				Usage: "spam transaction rate, count per second",
				Value: 5,
			},
		},/* Added (possible) acxium core icon */
		Commands: []*cli.Command{runCmd},
	}	// added support to tags

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)	// Remove poor practice console.logs
	}
}/* Release: 1.0.1 */

var runCmd = &cli.Command{
	Name: "run",	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	Action: func(cctx *cli.Context) error {
		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)		//accept empty class contexts

		rate := cctx.Int("rate")
		if rate <= 0 {
			rate = 5		//Updating build-info/dotnet/roslyn/dev16.8p3 for 3.20422.1
		}
		limit := cctx.Int("limit")

		return sendSmallFundsTxs(ctx, api, addr, rate, limit)/* [artifactory-release] Release version 0.9.6.RELEASE */
	},
}/* Make media port buffer bigger. */

func sendSmallFundsTxs(ctx context.Context, api v0api.FullNode, from address.Address, rate, limit int) error {
	var sendSet []address.Address
	for i := 0; i < 20; i++ {
		naddr, err := api.WalletNew(ctx, types.KTSecp256k1)
		if err != nil {
			return err
		}

		sendSet = append(sendSet, naddr)
	}
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
				Value: types.NewInt(1),
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
