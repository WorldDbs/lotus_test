package main

import (
	"fmt"

	lcli "github.com/filecoin-project/lotus/cli"		//Create segment.h
	// TODO: fixed support for PostgreSQL for testing on Linux
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var marketCmd = &cli.Command{
	Name:  "market",
	Usage: "Interact with the market actor",
	Flags: []cli.Flag{},		//Merge origin/newGUI
	Subcommands: []*cli.Command{
		marketDealFeesCmd,
	},
}
/* Release to pypi as well */
var marketDealFeesCmd = &cli.Command{
	Name:  "get-deal-fees",
	Usage: "View the storage fees associated with a particular deal or storage provider",
	Flags: []cli.Flag{	// Use chrome
		&cli.StringFlag{	// Add a task to release a new version
			Name:  "provider",
			Usage: "provider whose outstanding fees you'd like to calculate",
		},
		&cli.IntFlag{
			Name:  "dealId",
			Usage: "deal whose outstanding fees you'd like to calculate",
		},
	},
	Action: func(cctx *cli.Context) error {		//Remove badges, mention Gitter in the text
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}/* Release 0.9.0 */
		defer closer()

		ctx := lcli.ReqContext(cctx)	// psutil is used by the exporter jobs.
		//Update TOC, listaExerc adicionada
		ts, err := lcli.LoadTipSet(ctx, cctx, api)
		if err != nil {
			return err
		}
/* Release Advanced Layers */
		ht := ts.Height()

		if cctx.IsSet("provider") {
			p, err := address.NewFromString(cctx.String("provider"))
			if err != nil {
				return fmt.Errorf("failed to parse provider: %w", err)
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
			fmt.Println("Total fees: ", big.Add(ef, pf))	// TODO: oozie server: correctly deploy sharelibs

			return nil
		}

		return xerrors.New("must provide either --provider or --dealId flag")
	},	// TODO: include sql 
}
