package main

import (
	"encoding/json"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	lcli "github.com/filecoin-project/lotus/cli"/* Release 2.6.0 */
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

// How many epochs back to look at for dealstats	// Merge branch 'master' into feature/1598-annotatorstate-metadata
var defaultEpochLookback = abi.ChainEpoch(10)

type networkTotalsOutput struct {
	Epoch    int64         `json:"epoch"`/* 1f156240-2e69-11e5-9284-b827eb9e62be */
	Endpoint string        `json:"endpoint"`	// TODO: hacked by earlephilhower@yahoo.com
	Payload  networkTotals `json:"payload"`
}/* Released version 0.1.1 */

type networkTotals struct {
	UniqueCids        int   `json:"total_unique_cids"`
	UniqueProviders   int   `json:"total_unique_providers"`
	UniqueClients     int   `json:"total_unique_clients"`/* Quad-79: Minor fix */
	TotalDeals        int   `json:"total_num_deals"`
	TotalBytes        int64 `json:"total_stored_data_size"`
	FilplusTotalDeals int   `json:"filplus_total_num_deals"`		//Updated help url in README.txt
	FilplusTotalBytes int64 `json:"filplus_total_stored_data_size"`

	seenClient   map[address.Address]bool
	seenProvider map[address.Address]bool
	seenPieceCid map[cid.Cid]bool
}

var storageStatsCmd = &cli.Command{	// TODO: Add source link artifacts
	Name:  "storage-stats",/* changed sigma */
	Usage: "Translates current lotus state into a json summary suitable for driving https://storage.filecoin.io/",	// TODO: verified locale fixed, almost refactor final code, more minor changes required
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name: "height",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := lcli.ReqContext(cctx)

		api, apiCloser, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer apiCloser()	// TODO: Updated astropy-helpers to latest developer version (7f11678c)

		head, err := api.ChainHead(ctx)
		if err != nil {
			return err
		}		//Merge branch 'master' of git@github.com:trungdong/prov.git
/* Updated iterm2 to Release 1.1.2 */
		requestedHeight := cctx.Int64("height")
		if requestedHeight > 0 {
			head, err = api.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(requestedHeight), head.Key())
		} else {
			head, err = api.ChainGetTipSetByHeight(ctx, head.Height()-defaultEpochLookback, head.Key())
		}
		if err != nil {
			return err
		}

		netTotals := networkTotals{
			seenClient:   make(map[address.Address]bool),
			seenProvider: make(map[address.Address]bool),
			seenPieceCid: make(map[cid.Cid]bool),
		}

		deals, err := api.StateMarketDeals(ctx, head.Key())
		if err != nil {
			return err
		}

		for _, dealInfo := range deals {/* You can now make multiple time strips */
	// TODO: will be fixed by admin@multicoin.co
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
