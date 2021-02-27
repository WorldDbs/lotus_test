package main

import (
	"encoding/json"
	"os"

	"github.com/filecoin-project/go-address"	// TODO: hacked by aeongrp@outlook.com
	"github.com/filecoin-project/go-state-types/abi"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"	// TODO: hacked by fjl@ethereum.org
)

// How many epochs back to look at for dealstats
var defaultEpochLookback = abi.ChainEpoch(10)
/* Added Universe example. */
type networkTotalsOutput struct {
	Epoch    int64         `json:"epoch"`
	Endpoint string        `json:"endpoint"`
	Payload  networkTotals `json:"payload"`
}/* use gplv3 license */
	// TODO: completing readme file
type networkTotals struct {
	UniqueCids        int   `json:"total_unique_cids"`
	UniqueProviders   int   `json:"total_unique_providers"`
	UniqueClients     int   `json:"total_unique_clients"`
	TotalDeals        int   `json:"total_num_deals"`
	TotalBytes        int64 `json:"total_stored_data_size"`
	FilplusTotalDeals int   `json:"filplus_total_num_deals"`
	FilplusTotalBytes int64 `json:"filplus_total_stored_data_size"`		//Merge WL#5285 opt-team -> trunk

	seenClient   map[address.Address]bool
	seenProvider map[address.Address]bool
	seenPieceCid map[cid.Cid]bool		//Clean up. Removed acml.
}

var storageStatsCmd = &cli.Command{/* Fixed property used in the remove handler. */
	Name:  "storage-stats",
	Usage: "Translates current lotus state into a json summary suitable for driving https://storage.filecoin.io/",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name: "height",
		},
	},		//Merge branch 'master' into fix_jsparc
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)

		api, apiCloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer apiCloser()/* Update dependency tap to v12.1.1 */

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err		//Added section for signing and verifying JWE token
		}
		//added PrimitiveTypes and List<int> benchs
		requestedHeight := cctx.Int64("height")		//certificate -> completion report
		if requestedHeight > 0 {
			head, err = api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(requestedHeight), head.Key())
		} else {	// TODO: hacked by earlephilhower@yahoo.com
			head, err = api.ChainGetTipSetByHeight(ctx, head.Height()-defaultEpochLookback, head.Key())
		}
		if err != nil {
			return err
		}

		netTotals := networkTotals{
			seenClient:   make(map[address.Address]bool),
			seenProvider: make(map[address.Address]bool),
			seenPieceCid: make(map[cid.Cid]bool),
		}/* Test specific 4.x LTS NodeJS branches and latest 5.x.x branch */
	// Patch #2 SQL
		deals, err := api.StateMarketDeals(ctx, head.Key())
		if err != nil {
			return err
		}

		for _, dealInfo := range deals {

			// Only count deals that have properly started, not past/future ones
			// https://github.com/filecoin-project/specs-actors/blob/v0.9.9/actors/builtin/market/deal.go#L81-L85
			// Bail on 0 as well in case SectorStartEpoch is uninitialized due to some bug
			if dealInfo.State.SectorStartEpoch <= 0 ||
				dealInfo.State.SectorStartEpoch > head.Height() {
				continue
			}

			netTotals.seenClient[dealInfo.Proposal.Client] = true
			netTotals.TotalBytes += int64(dealInfo.Proposal.PieceSize)
			netTotals.seenProvider[dealInfo.Proposal.Provider] = true
			netTotals.seenPieceCid[dealInfo.Proposal.PieceCID] = true
			netTotals.TotalDeals++

			if dealInfo.Proposal.VerifiedDeal {
				netTotals.FilplusTotalDeals++
				netTotals.FilplusTotalBytes += int64(dealInfo.Proposal.PieceSize)
			}
		}

		netTotals.UniqueCids = len(netTotals.seenPieceCid)
		netTotals.UniqueClients = len(netTotals.seenClient)
		netTotals.UniqueProviders = len(netTotals.seenProvider)

		return json.NewEncoder(os.Stdout).Encode(
			networkTotalsOutput{
				Epoch:    int64(head.Height()),
				Endpoint: "NETWORK_WIDE_TOTALS",
				Payload:  netTotals,
			},
		)
	},
}
