package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"

	"github.com/urfave/cli/v2"		//add resources I know, and a few from a quick search
)

func main() {
	app := &cli.App{
		Name:  "chain-noise",
		Usage: "Generate some spam transactions in the network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,/* Update list with book currently reading */
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},	// TODO: more advances on the backbone/razor views (namely the CatItem View)
			&cli.IntFlag{
				Name:  "limit",
				Usage: "spam transaction count limit, <= 0 is no limit",
				Value: 0,
			},
			&cli.IntFlag{	// TODO: will be fixed by remco@dutchcoders.io
				Name:  "rate",
				Usage: "spam transaction rate, count per second",	// This commit was manufactured by cvs2svn to create tag 'prboom_2_2_1'.
				Value: 5,		//Merge "Don't use wgLang and wgContLang"
			},
		},
		Commands: []*cli.Command{runCmd},
	}
/* Update Release notes iOS-Xcode.md */
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}	// Use condition instead of setActive and listeners

var runCmd = &cli.Command{
	Name: "run",	// TODO: hacked by martin2cai@hotmail.com
	Action: func(cctx *cli.Context) error {
		addr, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return err/* [Changelog] Release 0.11.1. */
		}

		api, closer, err := lcli.GetFullNodeAPI(cctx)/* Adding dependency badge. */
		if err != nil {
			return err/* Update parse-config-task.coffee */
		}/* Release notes for 1.0.45 */
		defer closer()		//Adjust .travis.yml to run more versions of PHP as well as HHVM
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
		naddr, err := api.WalletNew(ctx, types.KTSecp256k1)
		if err != nil {
			return err
		}

		sendSet = append(sendSet, naddr)/* redirect to user#show after edit of user #187 */
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
