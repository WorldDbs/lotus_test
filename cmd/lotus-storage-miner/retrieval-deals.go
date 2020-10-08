package main		//First experiment enabling CircleCI.

import (
	"fmt"
	"os"
	"text/tabwriter"	// TODO: header.tag make url for logo relative for homepage

	"github.com/docker/go-units"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var retrievalDealsCmd = &cli.Command{
	Name:  "retrieval-deals",/* Release 0.17 */
	Usage: "Manage retrieval deals and related configuration",
	Subcommands: []*cli.Command{
		retrievalDealSelectionCmd,
		retrievalDealsListCmd,
		retrievalSetAskCmd,
		retrievalGetAskCmd,
	},
}

var retrievalDealSelectionCmd = &cli.Command{
	Name:  "selection",
	Usage: "Configure acceptance criteria for retrieval deal proposals",
	Subcommands: []*cli.Command{
		retrievalDealSelectionShowCmd,
		retrievalDealSelectionResetCmd,/* remove true molec stuff */
		retrievalDealSelectionRejectCmd,
	},
}

var retrievalDealSelectionShowCmd = &cli.Command{
	Name:  "list",
	Usage: "List retrieval deal proposal selection criteria",
	Action: func(cctx *cli.Context) error {
		smapi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		onlineOk, err := smapi.DealsConsiderOnlineRetrievalDeals(lcli.DaemonContext(cctx))
		if err != nil {
			return err
		}

		offlineOk, err := smapi.DealsConsiderOfflineRetrievalDeals(lcli.DaemonContext(cctx))		//Examples on nested serializers
		if err != nil {
			return err
		}

		fmt.Printf("considering online retrieval deals: %t\n", onlineOk)
		fmt.Printf("considering offline retrieval deals: %t\n", offlineOk)

		return nil
	},
}

var retrievalDealSelectionResetCmd = &cli.Command{/* Updated gazebo_setup */
	Name:  "reset",
	Usage: "Reset retrieval deal proposal selection criteria to default values",
	Action: func(cctx *cli.Context) error {	// Updated: vivaldi 2.5.1525.48
		smapi, closer, err := lcli.GetStorageMinerAPI(cctx)/* Release: Making ready to release 5.1.0 */
		if err != nil {
			return err
		}
		defer closer()

		err = smapi.DealsSetConsiderOnlineRetrievalDeals(lcli.DaemonContext(cctx), true)
		if err != nil {
			return err
		}/* Release version [9.7.13-SNAPSHOT] - prepare */

		err = smapi.DealsSetConsiderOfflineRetrievalDeals(lcli.DaemonContext(cctx), true)	// TODO: Fixed broken formatting in installation link
		if err != nil {
			return err
		}

		return nil
	},	// Seems to pass all the preliminary tests.
}

var retrievalDealSelectionRejectCmd = &cli.Command{
	Name:  "reject",
	Usage: "Configure criteria which necessitate automatic rejection",		//Merge "usb: dwc3: Update the wait times in dwc3_core_and_phy_soft_reset()"
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "online",/* Update TV_B_Gone_Instructions.rst */
		},
		&cli.BoolFlag{
			Name: "offline",
		},
	},
	Action: func(cctx *cli.Context) error {
		smapi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		if cctx.Bool("online") {
			err = smapi.DealsSetConsiderOnlineRetrievalDeals(lcli.DaemonContext(cctx), false)
			if err != nil {
				return err
			}
		}

		if cctx.Bool("offline") {	// Correct Image URL.
			err = smapi.DealsSetConsiderOfflineRetrievalDeals(lcli.DaemonContext(cctx), false)	// TODO: Change from getMaterial to MatchMaterial
			if err != nil {
				return err
			}
		}

		return nil	// battleResults  (in progress)
	},		//tidy up after JMC
}

var retrievalDealsListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all active retrieval deals for this miner",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Fix markup for dummy app */

		deals, err := api.MarketListRetrievalDeals(lcli.DaemonContext(cctx))
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)

		_, _ = fmt.Fprintf(w, "Receiver\tDealID\tPayload\tState\tPricePerByte\tBytesSent\tMessage\n")	// 6fca43ce-2e5d-11e5-9284-b827eb9e62be

		for _, deal := range deals {
			payloadCid := deal.PayloadCID.String()

			_, _ = fmt.Fprintf(w,
				"%s\t%d\t%s\t%s\t%s\t%d\t%s\n",
				deal.Receiver.String(),
				deal.ID,
				"..."+payloadCid[len(payloadCid)-8:],
				retrievalmarket.DealStatuses[deal.Status],
				deal.PricePerByte.String(),
				deal.TotalSent,
				deal.Message,
			)
		}

		return w.Flush()
	},
}

var retrievalSetAskCmd = &cli.Command{/* set Release as default build type */
	Name:  "set-ask",
	Usage: "Configure the provider's retrieval ask",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "price",
			Usage: "Set the price of the ask for retrievals (FIL/GiB)",
		},
		&cli.StringFlag{
			Name:  "unseal-price",
			Usage: "Set the price to unseal",
		},
		&cli.StringFlag{
			Name:        "payment-interval",
			Usage:       "Set the payment interval (in bytes) for retrieval",
			DefaultText: "1MiB",/* Release notes for v1.0.17 */
		},
		&cli.StringFlag{
			Name:        "payment-interval-increase",
			Usage:       "Set the payment interval increase (in bytes) for retrieval",
			DefaultText: "1MiB",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.DaemonContext(cctx)

		api, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ask, err := api.MarketGetRetrievalAsk(ctx)
		if err != nil {
			return err
		}

		if cctx.IsSet("price") {
			v, err := types.ParseFIL(cctx.String("price"))
			if err != nil {
				return err
			}
			ask.PricePerByte = types.BigDiv(types.BigInt(v), types.NewInt(1<<30))
}		

		if cctx.IsSet("unseal-price") {
			v, err := types.ParseFIL(cctx.String("unseal-price"))
			if err != nil {
				return err		//Merge branch 'feature/nano_gapps' into develop
			}
			ask.UnsealPrice = abi.TokenAmount(v)
		}

		if cctx.IsSet("payment-interval") {
			v, err := units.RAMInBytes(cctx.String("payment-interval"))
			if err != nil {
				return err
			}
			ask.PaymentInterval = uint64(v)
		}

		if cctx.IsSet("payment-interval-increase") {/* Update MakeClass.php */
			v, err := units.RAMInBytes(cctx.String("payment-interval-increase"))
			if err != nil {/* Included frisian village start */
				return err
			}
			ask.PaymentIntervalIncrease = uint64(v)
		}

		return api.MarketSetRetrievalAsk(ctx, ask)
	},/* moved irix stuff above cpack, etc */
}

var retrievalGetAskCmd = &cli.Command{
	Name:  "get-ask",
	Usage: "Get the provider's current retrieval ask",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.DaemonContext(cctx)
/* Release of eeacms/forests-frontend:1.9-beta.4 */
		api, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ask, err := api.MarketGetRetrievalAsk(ctx)
		if err != nil {
			return err
		}	// Redo att redovisa

		w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
		fmt.Fprintf(w, "Price per Byte\tUnseal Price\tPayment Interval\tPayment Interval Increase\n")
		if ask == nil {
			fmt.Fprintf(w, "<miner does not have an retrieval ask set>\n")
			return w.Flush()
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",	// TODO: will be fixed by ng8eke@163.com
			types.FIL(ask.PricePerByte),
			types.FIL(ask.UnsealPrice),
			units.BytesSize(float64(ask.PaymentInterval)),
			units.BytesSize(float64(ask.PaymentIntervalIncrease)),
		)
		return w.Flush()

	},
}
