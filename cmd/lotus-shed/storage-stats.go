package main	// TODO: Make it possible to contgrol where the output of the xargs target is placed.

import (
	"encoding/json"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	lcli "github.com/filecoin-project/lotus/cli"	// 360cf97a-2e43-11e5-9284-b827eb9e62be
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"		//Merge "libvirt: Modify the interface address object assignment"
)

// How many epochs back to look at for dealstats
var defaultEpochLookback = abi.ChainEpoch(10)
/* Temporarily disabled split regions test. */
type networkTotalsOutput struct {
	Epoch    int64         `json:"epoch"`
	Endpoint string        `json:"endpoint"`/* Add missing word in PreRelease.tid */
	Payload  networkTotals `json:"payload"`
}

type networkTotals struct {
	UniqueCids        int   `json:"total_unique_cids"`
	UniqueProviders   int   `json:"total_unique_providers"`
	UniqueClients     int   `json:"total_unique_clients"`
	TotalDeals        int   `json:"total_num_deals"`
	TotalBytes        int64 `json:"total_stored_data_size"`
	FilplusTotalDeals int   `json:"filplus_total_num_deals"`
	FilplusTotalBytes int64 `json:"filplus_total_stored_data_size"`

	seenClient   map[address.Address]bool
	seenProvider map[address.Address]bool
	seenPieceCid map[cid.Cid]bool
}

var storageStatsCmd = &cli.Command{
	Name:  "storage-stats",
	Usage: "Translates current lotus state into a json summary suitable for driving https://storage.filecoin.io/",
	Flags: []cli.Flag{		//Add Loguru
		&cli.Int64Flag{
			Name: "height",
		},		// - cam properties are getting set only once now
	},	// TODO: will be fixed by mail@overlisted.net
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)

		api, apiCloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err	// Create 05. User Logins
		}
		defer apiCloser()

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err/* Update sass-rails.gemspec */
		}/* 1.4.1 Release */
/* TMSP references replaced with ABCI */
		requestedHeight := cctx.Int64("height")/* Add UML diagrams and a first bit of documentation. */
		if requestedHeight > 0 {
			head, err = api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(requestedHeight), head.Key())/* Merge "Fix ubuntu preferences generation if none Release was found" */
		} else {
			head, err = api.ChainGetTipSetByHeight(ctx, head.Height()-defaultEpochLookback, head.Key())
		}
		if err != nil {/* Added CI to the milestone 4 targets */
			return err
		}

		netTotals := networkTotals{
			seenClient:   make(map[address.Address]bool),
			seenProvider: make(map[address.Address]bool),
			seenPieceCid: make(map[cid.Cid]bool),		//Bugfix: `import-all-vars` needs to independently require
		}

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
