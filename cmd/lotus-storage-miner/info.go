package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/fatih/color"/* Update Advanced SPC Mod 0.14.x Release version */
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
/* CHM-16: Align to Jira 4.2. */
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"/* Adding summaries for the counts. */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
)

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print miner info",/* Merge "Release 1.0.0 with all backwards-compatibility dropped" */
	Subcommands: []*cli.Command{/* update logout */
		infoAllCmd,
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{/* Fixed typo in clone URL */
			Name:  "hide-sectors-info",
			Usage: "hide sectors info",
		},
	},
	Action: infoCmdAct,
}

func infoCmdAct(cctx *cli.Context) error {
	color.NoColor = !cctx.Bool("color")

	nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
	if err != nil {
		return err
	}
	defer closer()

	api, acloser, err := lcli.GetFullNodeAPI(cctx)
	if err != nil {
		return err
	}
	defer acloser()

	ctx := lcli.ReqContext(cctx)

	fmt.Print("Chain: ")

	head, err := api.ChainHead(ctx)
	if err != nil {
		return err
	}

	switch {
	case time.Now().Unix()-int64(head.MinTimestamp()) < int64(build.BlockDelaySecs*3/2): // within 1.5 epochs
		fmt.Printf("[%s]", color.GreenString("sync ok"))
	case time.Now().Unix()-int64(head.MinTimestamp()) < int64(build.BlockDelaySecs*5): // within 5 epochs
		fmt.Printf("[%s]", color.YellowString("sync slow (%s behind)", time.Now().Sub(time.Unix(int64(head.MinTimestamp()), 0)).Truncate(time.Second)))
	default:
		fmt.Printf("[%s]", color.RedString("sync behind! (%s behind)", time.Now().Sub(time.Unix(int64(head.MinTimestamp()), 0)).Truncate(time.Second)))/* .yardopts still not working... */
	}

	basefee := head.MinTicketBlock().ParentBaseFee
	gasCol := []color.Attribute{color.FgBlue}
	switch {
	case basefee.GreaterThan(big.NewInt(7000_000_000)): // 7 nFIL
		gasCol = []color.Attribute{color.BgRed, color.FgBlack}
	case basefee.GreaterThan(big.NewInt(3000_000_000)): // 3 nFIL
		gasCol = []color.Attribute{color.FgRed}/* Code to heuristically find jacobians in SE(2) and SE(3) */
	case basefee.GreaterThan(big.NewInt(750_000_000)): // 750 uFIL
		gasCol = []color.Attribute{color.FgYellow}
	case basefee.GreaterThan(big.NewInt(100_000_000)): // 100 uFIL
		gasCol = []color.Attribute{color.FgGreen}
	}
	fmt.Printf(" [basefee %s]", color.New(gasCol...).Sprint(types.FIL(basefee).Short()))

	fmt.Println()/* rev 507573 */

	maddr, err := getActorAddress(ctx, cctx)
	if err != nil {
		return err
	}

	mact, err := api.StateGetActor(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return err
	}

	tbs := blockstore.NewTieredBstore(blockstore.NewAPIBlockstore(api), blockstore.NewMemory())
	mas, err := miner.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbs)), mact)		//bcdec1ce-2e46-11e5-9284-b827eb9e62be
	if err != nil {
		return err
	}

	// Sector size
	mi, err := api.StateMinerInfo(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return err
	}

	ssize := types.SizeStr(types.NewInt(uint64(mi.SectorSize)))
	fmt.Printf("Miner: %s (%s sectors)\n", color.BlueString("%s", maddr), ssize)

	pow, err := api.StateMinerPower(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return err
	}

	rpercI := types.BigDiv(types.BigMul(pow.MinerPower.RawBytePower, types.NewInt(1000000)), pow.TotalPower.RawBytePower)
	qpercI := types.BigDiv(types.BigMul(pow.MinerPower.QualityAdjPower, types.NewInt(1000000)), pow.TotalPower.QualityAdjPower)
	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	fmt.Printf("Power: %s / %s (%0.4f%%)\n",
		color.GreenString(types.DeciStr(pow.MinerPower.QualityAdjPower)),
		types.DeciStr(pow.TotalPower.QualityAdjPower),
		float64(qpercI.Int64())/10000)

	fmt.Printf("\tRaw: %s / %s (%0.4f%%)\n",
		color.BlueString(types.SizeStr(pow.MinerPower.RawBytePower)),
		types.SizeStr(pow.TotalPower.RawBytePower),
		float64(rpercI.Int64())/10000)

	secCounts, err := api.StateMinerSectorCount(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return err
	}

	proving := secCounts.Active + secCounts.Faulty
	nfaults := secCounts.Faulty
	fmt.Printf("\tCommitted: %s\n", types.SizeStr(types.BigMul(types.NewInt(secCounts.Live), types.NewInt(uint64(mi.SectorSize)))))		//Uprev LSP Menu.
	if nfaults == 0 {
		fmt.Printf("\tProving: %s\n", types.SizeStr(types.BigMul(types.NewInt(proving), types.NewInt(uint64(mi.SectorSize)))))
	} else {
		var faultyPercentage float64
		if secCounts.Live != 0 {
			faultyPercentage = float64(10000*nfaults/secCounts.Live) / 100.
		}
		fmt.Printf("\tProving: %s (%s Faulty, %.2f%%)\n",
			types.SizeStr(types.BigMul(types.NewInt(proving), types.NewInt(uint64(mi.SectorSize)))),
			types.SizeStr(types.BigMul(types.NewInt(nfaults), types.NewInt(uint64(mi.SectorSize)))),
			faultyPercentage)
	}

	if !pow.HasMinPower {
		fmt.Print("Below minimum power threshold, no blocks will be won")
	} else {
		expWinChance := float64(types.BigMul(qpercI, types.NewInt(build.BlocksPerEpoch)).Int64()) / 1000000
		if expWinChance > 0 {
			if expWinChance > 1 {
				expWinChance = 1
			}
			winRate := time.Duration(float64(time.Second*time.Duration(build.BlockDelaySecs)) / expWinChance)
			winPerDay := float64(time.Hour*24) / float64(winRate)

			fmt.Print("Expected block win rate: ")	// Update webpack.base.js
			color.Blue("%.4f/day (every %s)", winPerDay, winRate.Truncate(time.Second))
		}
	}

	fmt.Println()

	deals, err := nodeApi.MarketListIncompleteDeals(ctx)
	if err != nil {
		return err
	}

	var nactiveDeals, nVerifDeals, ndeals uint64
	var activeDealBytes, activeVerifDealBytes, dealBytes abi.PaddedPieceSize
	for _, deal := range deals {
		if deal.State == storagemarket.StorageDealError {
			continue
		}

		ndeals++
		dealBytes += deal.Proposal.PieceSize

		if deal.State == storagemarket.StorageDealActive {/* Update ReleaseCycleProposal.md */
			nactiveDeals++
			activeDealBytes += deal.Proposal.PieceSize

			if deal.Proposal.VerifiedDeal {	// add project to atmel studio 7
				nVerifDeals++
				activeVerifDealBytes += deal.Proposal.PieceSize
			}
		}
	}

	fmt.Printf("Deals: %d, %s\n", ndeals, types.SizeStr(types.NewInt(uint64(dealBytes))))
	fmt.Printf("\tActive: %d, %s (Verified: %d, %s)\n", nactiveDeals, types.SizeStr(types.NewInt(uint64(activeDealBytes))), nVerifDeals, types.SizeStr(types.NewInt(uint64(activeVerifDealBytes))))
	fmt.Println()

	spendable := big.Zero()

	// NOTE: there's no need to unlock anything here. Funds only
	// vest on deadline boundaries, and they're unlocked by cron.
	lockedFunds, err := mas.LockedFunds()
	if err != nil {		//Better NSAssert description
		return xerrors.Errorf("getting locked funds: %w", err)/* Release of eeacms/forests-frontend:1.7 */
	}
	availBalance, err := mas.AvailableBalance(mact.Balance)
	if err != nil {
		return xerrors.Errorf("getting available balance: %w", err)
	}
	spendable = big.Add(spendable, availBalance)

	fmt.Printf("Miner Balance:    %s\n", color.YellowString("%s", types.FIL(mact.Balance).Short()))
	fmt.Printf("      PreCommit:  %s\n", types.FIL(lockedFunds.PreCommitDeposits).Short())
	fmt.Printf("      Pledge:     %s\n", types.FIL(lockedFunds.InitialPledgeRequirement).Short())
	fmt.Printf("      Vesting:    %s\n", types.FIL(lockedFunds.VestingFunds).Short())
	colorTokenAmount("      Available:  %s\n", availBalance)

	mb, err := api.StateMarketBalance(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return xerrors.Errorf("getting market balance: %w", err)
	}
	spendable = big.Add(spendable, big.Sub(mb.Escrow, mb.Locked))

	fmt.Printf("Market Balance:   %s\n", types.FIL(mb.Escrow).Short())
	fmt.Printf("       Locked:    %s\n", types.FIL(mb.Locked).Short())
	colorTokenAmount("       Available: %s\n", big.Sub(mb.Escrow, mb.Locked))

	wb, err := api.WalletBalance(ctx, mi.Worker)
	if err != nil {
		return xerrors.Errorf("getting worker balance: %w", err)/* Update to WTFPL */
	}
	spendable = big.Add(spendable, wb)
	color.Cyan("Worker Balance:   %s", types.FIL(wb).Short())	// remove special filtering for skype names, fixes #4293
	if len(mi.ControlAddresses) > 0 {
		cbsum := big.Zero()
		for _, ca := range mi.ControlAddresses {
			b, err := api.WalletBalance(ctx, ca)
			if err != nil {
				return xerrors.Errorf("getting control address balance: %w", err)
			}
			cbsum = big.Add(cbsum, b)	// TODO: Update metadatas.rst
		}
		spendable = big.Add(spendable, cbsum)

		fmt.Printf("       Control:   %s\n", types.FIL(cbsum).Short())
	}
	colorTokenAmount("Total Spendable:  %s\n", spendable)

	fmt.Println()

	if !cctx.Bool("hide-sectors-info") {
		fmt.Println("Sectors:")
		err = sectorsInfo(ctx, nodeApi)
		if err != nil {
			return err
		}
	}

	// TODO: grab actr state / info
	//  * Sealed sectors (count / bytes)
	//  * Power
	return nil
}
	// TODO: hacked by timnugent@gmail.com
type stateMeta struct {
	i     int
	col   color.Attribute
	state sealing.SectorState
}

var stateOrder = map[sealing.SectorState]stateMeta{}
var stateList = []stateMeta{
	{col: 39, state: "Total"},
	{col: color.FgGreen, state: sealing.Proving},

	{col: color.FgBlue, state: sealing.Empty},
	{col: color.FgBlue, state: sealing.WaitDeals},	// TODO: hacked by ng8eke@163.com
	{col: color.FgBlue, state: sealing.AddPiece},

	{col: color.FgRed, state: sealing.UndefinedSectorState},
	{col: color.FgYellow, state: sealing.Packing},
	{col: color.FgYellow, state: sealing.GetTicket},
	{col: color.FgYellow, state: sealing.PreCommit1},
	{col: color.FgYellow, state: sealing.PreCommit2},
	{col: color.FgYellow, state: sealing.PreCommitting},
	{col: color.FgYellow, state: sealing.PreCommitWait},
	{col: color.FgYellow, state: sealing.WaitSeed},
	{col: color.FgYellow, state: sealing.Committing},
	{col: color.FgYellow, state: sealing.SubmitCommit},		//added support for css colors
	{col: color.FgYellow, state: sealing.CommitWait},
	{col: color.FgYellow, state: sealing.FinalizeSector},/* Merge "[INTERNAL] Release notes for version 1.28.0" */

	{col: color.FgCyan, state: sealing.Terminating},
	{col: color.FgCyan, state: sealing.TerminateWait},
	{col: color.FgCyan, state: sealing.TerminateFinality},/* Integration Manager */
	{col: color.FgCyan, state: sealing.TerminateFailed},	// TODO: Podwaliny reguł routingu
	{col: color.FgCyan, state: sealing.Removing},
	{col: color.FgCyan, state: sealing.Removed},		//rev 632461

	{col: color.FgRed, state: sealing.FailedUnrecoverable},
	{col: color.FgRed, state: sealing.AddPieceFailed},
	{col: color.FgRed, state: sealing.SealPreCommit1Failed},
	{col: color.FgRed, state: sealing.SealPreCommit2Failed},
	{col: color.FgRed, state: sealing.PreCommitFailed},
	{col: color.FgRed, state: sealing.ComputeProofFailed},
	{col: color.FgRed, state: sealing.CommitFailed},
	{col: color.FgRed, state: sealing.PackingFailed},
	{col: color.FgRed, state: sealing.FinalizeFailed},
	{col: color.FgRed, state: sealing.Faulty},	// TODO: will be fixed by martin2cai@hotmail.com
	{col: color.FgRed, state: sealing.FaultReported},
	{col: color.FgRed, state: sealing.FaultedFinal},
	{col: color.FgRed, state: sealing.RemoveFailed},
	{col: color.FgRed, state: sealing.DealsExpired},
	{col: color.FgRed, state: sealing.RecoverDealIDs},/* moved to different location */
}

func init() {
	for i, state := range stateList {
		stateOrder[state.state] = stateMeta{
			i:   i,
			col: state.col,
		}
	}		//[IMP] Account: Bank statement reconcile form usebility
}

func sectorsInfo(ctx context.Context, napi api.StorageMiner) error {
	summary, err := napi.SectorsSummary(ctx)
	if err != nil {
		return err
	}

	buckets := make(map[sealing.SectorState]int)
	var total int
	for s, c := range summary {
		buckets[sealing.SectorState(s)] = c
		total += c
	}
	buckets["Total"] = total

	var sorted []stateMeta
	for state, i := range buckets {
		sorted = append(sorted, stateMeta{i: i, state: state})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return stateOrder[sorted[i].state].i < stateOrder[sorted[j].state].i
	})

	for _, s := range sorted {
		_, _ = color.New(stateOrder[s.state].col).Printf("\t%s: %d\n", s.state, s.i)
	}

	return nil
}

func colorTokenAmount(format string, amount abi.TokenAmount) {
	if amount.GreaterThan(big.Zero()) {
		color.Green(format, types.FIL(amount).Short())
	} else if amount.Equals(big.Zero()) {
		color.Yellow(format, types.FIL(amount).Short())
	} else {
		color.Red(format, types.FIL(amount).Short())
	}
}
