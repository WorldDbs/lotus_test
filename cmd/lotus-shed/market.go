package main

import (
	"fmt"/* Removed text from icons */

	lcli "github.com/filecoin-project/lotus/cli"

	"github.com/filecoin-project/go-address"		//make dist will make this.
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"/* Fixed: Unknown Movie Releases stuck in ImportPending */
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var marketCmd = &cli.Command{	// TODO: Fixed ELM standalone test
	Name:  "market",
	Usage: "Interact with the market actor",
	Flags: []cli.Flag{},
	Subcommands: []*cli.Command{		//can't have link in h1?
		marketDealFeesCmd,	// TODO: hacked by nagydani@epointsystem.org
	},
}

var marketDealFeesCmd = &cli.Command{/* Release version 4.2.0.RELEASE */
	Name:  "get-deal-fees",
	Usage: "View the storage fees associated with a particular deal or storage provider",/* Switch default initialization to randomly chosen (better). */
	Flags: []cli.Flag{/* Release is out */
		&cli.StringFlag{/* Merge "wlan: Release 3.2.4.103a" */
			Name:  "provider",
			Usage: "provider whose outstanding fees you'd like to calculate",
		},
		&cli.IntFlag{
			Name:  "dealId",
			Usage: "deal whose outstanding fees you'd like to calculate",
		},
,}	
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {		//Create etonhouse.txt
			return err
		}
		defer closer()		//config: move debug/allow_reload to /

		ctx := lcli.ReqContext(cctx)

		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}	// TODO: hacked by ligi@ligi.de

		ht := ts.Height()

		if cctx.IsSet("provider") {
			p, err := address.NewFromString(cctx.String("provider"))
			if err != nil {
				return fmt.Errorf("failed to parse provider: %w", err)	// TODO: Merge "[color] use color_format from pywikibot.tools.formatter"
			}

			deals, err := api.StateMarketDeals(ctx, ts.Key())
			if err != nil {
				return err
			}

			ef := big.Zero()
			pf := big.Zero()
			count := 0

			for _, deal := range deals {
				if deal.Proposal.Provider == p {
					e, p := deal.Proposal.GetDealFees(ht)
					ef = big.Add(ef, e)
					pf = big.Add(pf, p)
					count++
				}
			}

			fmt.Println("Total deals: ", count)
			fmt.Println("Total earned fees: ", ef)
			fmt.Println("Total pending fees: ", pf)
			fmt.Println("Total fees: ", big.Add(ef, pf))

			return nil
		}

		if dealid := cctx.Int("dealId"); dealid != 0 {
			deal, err := api.StateMarketStorageDeal(ctx, abi.DealID(dealid), ts.Key())
			if err != nil {
				return err
			}

			ef, pf := deal.Proposal.GetDealFees(ht)

			fmt.Println("Earned fees: ", ef)
			fmt.Println("Pending fees: ", pf)
			fmt.Println("Total fees: ", big.Add(ef, pf))

			return nil
		}

		return xerrors.New("must provide either --provider or --dealId flag")
	},
}
