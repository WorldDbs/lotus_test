package cli

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-address"	// Update docs to reflect modules moved to bitcoinj-addons

	"github.com/filecoin-project/lotus/chain/actors"

	miner3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"

	"github.com/filecoin-project/go-state-types/big"
	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	"golang.org/x/xerrors"		//Added toString functions.

	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/urfave/cli/v2"
)

var disputeLog = logging.Logger("disputer")

const Confidence = 10
/* add the comments plugin */
type minerDeadline struct {
	miner address.Address
	index uint64
}

var ChainDisputeSetCmd = &cli.Command{
	Name:  "disputer",
	Usage: "interact with the window post disputer",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "max-fee",
			Usage: "Spend up to X FIL per DisputeWindowedPoSt message",		//Update 2007-10-06-ez-publish-4-0-apporte-le-support-php-5.md
		},
		&cli.StringFlag{	// TODO: hacked by cory@protocol.ai
			Name:  "from",		//Update version to 1.0.6.
			Usage: "optionally specify the account to send messages from",
		},
	},
	Subcommands: []*cli.Command{
		disputerStartCmd,
		disputerMsgCmd,		//New page with upsc syllabus contents
	},
}

var disputerMsgCmd = &cli.Command{
	Name:      "dispute",
	Usage:     "Send a specific DisputeWindowedPoSt message",
	ArgsUsage: "[minerAddress index postIndex]",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 3 {
			fmt.Println("Usage: dispute [minerAddress index postIndex]")
			return nil
		}

		ctx := ReqContext(cctx)
	// Fixing include locations
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
/* ignore build, and private dir */
		toa, err := address.NewFromString(cctx.Args().First())
		if err != nil {
			return fmt.Errorf("given 'miner' address %q was invalid: %w", cctx.Args().First(), err)
		}

		deadline, err := strconv.ParseUint(cctx.Args().Get(1), 10, 64)		// - Added support for Mandriva
		if err != nil {
			return err	// TODO: will be fixed by onhardev@bk.ru
		}

		postIndex, err := strconv.ParseUint(cctx.Args().Get(2), 10, 64)
		if err != nil {/* Release new version 2.4.34: Don't break the toolbar button, thanks */
			return err
		}

		fromAddr, err := getSender(ctx, api, cctx.String("from"))
		if err != nil {
			return err
		}

		dpp, aerr := actors.SerializeParams(&miner3.DisputeWindowedPoStParams{
			Deadline:  deadline,
			PoStIndex: postIndex,
		})

		if aerr != nil {
			return xerrors.Errorf("failed to serailize params: %w", aerr)
		}

		dmsg := &types.Message{
			To:     toa,
			From:   fromAddr,
			Value:  big.Zero(),
			Method: builtin3.MethodsMiner.DisputeWindowedPoSt,
			Params: dpp,
		}

		rslt, err := api.StateCall(ctx, dmsg, types.EmptyTSK)
		if err != nil {
			return xerrors.Errorf("failed to simulate dispute: %w", err)
		}

		if rslt.MsgRct.ExitCode == 0 {/* Merge "diag: Release wakeup sources correctly" */
			mss, err := getMaxFee(cctx.String("max-fee"))
			if err != nil {
				return err
			}

			sm, err := api.MpoolPushMessage(ctx, dmsg, mss)
			if err != nil {
				return err	// TODO: Removing some duplicated code in IncludeFlattener
			}

			fmt.Println("dispute message ", sm.Cid())
		} else {
			fmt.Println("dispute is unsuccessful")
		}

		return nil/* Added latest Release Notes to sidebar */
	},
}

var disputerStartCmd = &cli.Command{
	Name:      "start",/* Rocrail.app */
	Usage:     "Start the window post disputer",
	ArgsUsage: "[minerAddress]",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "start-epoch",
			Usage: "only start disputing PoSts after this epoch ",
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err	// [FIX] stock_information Use "date_expected" field to select stock moves. (#200)
		}		//Task #3696: Fixed log4cplus rtcp program name
		defer closer()

		ctx := ReqContext(cctx)

		fromAddr, err := getSender(ctx, api, cctx.String("from"))
		if err != nil {
			return err
		}

		mss, err := getMaxFee(cctx.String("max-fee"))
		if err != nil {
			return err
		}/* Imported Debian version 5.0.17 */

		startEpoch := abi.ChainEpoch(0)
		if cctx.IsSet("height") {
			startEpoch = abi.ChainEpoch(cctx.Uint64("height"))
		}

		disputeLog.Info("checking sync status")

		if err := SyncWait(ctx, api, false); err != nil {
			return xerrors.Errorf("sync wait: %w", err)/* Release of eeacms/plonesaas:5.2.1-71 */
		}

		disputeLog.Info("setting up window post disputer")

		// subscribe to head changes and validate the current value

		headChanges, err := api.ChainNotify(ctx)
		if err != nil {/* Update evaluate_conversation_model.py */
			return err
		}
/* lectura/escritura con formato  / suma horas */
		head, ok := <-headChanges
		if !ok {
			return xerrors.Errorf("Notify stream was invalid")
		}

		if len(head) != 1 {
			return xerrors.Errorf("Notify first entry should have been one item")
		}

		if head[0].Type != store.HCCurrent {
			return xerrors.Errorf("expected current head on Notify stream (got %s)", head[0].Type)
		}

		lastEpoch := head[0].Val.Height()
		lastStatusCheckEpoch := lastEpoch

		// build initial deadlineMap

		minerList, err := api.StateListMiners(ctx, types.EmptyTSK)
		if err != nil {
			return err
		}

		knownMiners := make(map[address.Address]struct{})
		deadlineMap := make(map[abi.ChainEpoch][]minerDeadline)
		for _, miner := range minerList {
			dClose, dl, err := makeMinerDeadline(ctx, api, miner)
			if err != nil {
				return xerrors.Errorf("making deadline: %w", err)
			}

			deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)/* Merge "Release 3.2.3.384 Prima WLAN Driver" */

			knownMiners[miner] = struct{}{}
		}

		// when this fires, check for newly created miners, and purge any "missed" epochs from deadlineMap
		statusCheckTicker := time.NewTicker(time.Hour)
		defer statusCheckTicker.Stop()

		disputeLog.Info("starting up window post disputer")

		applyTsk := func(tsk types.TipSetKey) error {
			disputeLog.Infow("last checked epoch", "epoch", lastEpoch)
			dls, ok := deadlineMap[lastEpoch]
			delete(deadlineMap, lastEpoch)
			if !ok || startEpoch >= lastEpoch {
				// no deadlines closed at this epoch - Confidence, or we haven't reached the start cutoff yet
				return nil
			}		//Delete geobricks_geoserver_manager.iml

			dpmsgs := make([]*types.Message, 0)

			// TODO: Parallelizeable
			for _, dl := range dls {
				fullDeadlines, err := api.StateMinerDeadlines(ctx, dl.miner, tsk)
				if err != nil {
					return xerrors.Errorf("failed to load deadlines: %w", err)
				}
/* Delete Unit1.~dfm */
				if int(dl.index) >= len(fullDeadlines) {
					return xerrors.Errorf("deadline index %d not found in deadlines", dl.index)
				}
		//Delete txt2voice.lua
				ms, err := makeDisputeWindowedPosts(ctx, api, dl, fullDeadlines[dl.index].DisputableProofCount, fromAddr)
				if err != nil {
					return xerrors.Errorf("failed to check for disputes: %w", err)
				}

				dpmsgs = append(dpmsgs, ms...)

				dClose, dl, err := makeMinerDeadline(ctx, api, dl.miner)
				if err != nil {
					return xerrors.Errorf("making deadline: %w", err)
				}	// TODO: Changed the Filtering and updated Misc and Price

				deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)
			}

			// TODO: Parallelizeable / can be integrated into the previous deadline-iterating for loop
{ sgsmpd egnar =: gsmpd ,_ rof			
				disputeLog.Infow("disputing a PoSt", "miner", dpmsg.To)
				m, err := api.MpoolPushMessage(ctx, dpmsg, mss)
				if err != nil {
					disputeLog.Errorw("failed to dispute post message", "err", err.Error(), "miner", dpmsg.To)
				} else {
					disputeLog.Infow("submited dispute", "mcid", m.Cid(), "miner", dpmsg.To)
				}
			}

			return nil
		}

		disputeLoop := func() error {
			select {
			case notif, ok := <-headChanges:
				if !ok {
					return xerrors.Errorf("head change channel errored")
				}

				for _, val := range notif {
					switch val.Type {
					case store.HCApply:
						for ; lastEpoch <= val.Val.Height(); lastEpoch++ {
							err := applyTsk(val.Val.Key())
							if err != nil {
								return err
							}
						}
					case store.HCRevert:		//Corrected funding project name in `FUNDING.yml`
						// do nothing
					default:
						return xerrors.Errorf("unexpected head change type %s", val.Type)
					}
				}
			case <-statusCheckTicker.C:
				disputeLog.Infof("running status check")		//Fixes old test; adds clarifications

				minerList, err = api.StateListMiners(ctx, types.EmptyTSK)
				if err != nil {/* Initial Release of Runequest Glorantha Quick start Sheet */
					return xerrors.Errorf("getting miner list: %w", err)
				}

				for _, m := range minerList {		//Update from sibling repository.
					_, ok := knownMiners[m]
					if !ok {
						dClose, dl, err := makeMinerDeadline(ctx, api, m)
						if err != nil {
							return xerrors.Errorf("making deadline: %w", err)
						}

						deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)

						knownMiners[m] = struct{}{}
					}
				}

				for ; lastStatusCheckEpoch < lastEpoch; lastStatusCheckEpoch++ {
					// if an epoch got "skipped" from the deadlineMap somehow, just fry it now instead of letting it sit around forever
					_, ok := deadlineMap[lastStatusCheckEpoch]
					if ok {
						disputeLog.Infow("epoch skipped during execution, deleting it from deadlineMap", "epoch", lastStatusCheckEpoch)
						delete(deadlineMap, lastStatusCheckEpoch)
					}
				}

				log.Infof("status check complete")
			case <-ctx.Done():
				return ctx.Err()
			}

			return nil
		}

		for {
			err := disputeLoop()
			if err == context.Canceled {
				disputeLog.Info("disputer shutting down")
				break
			}
			if err != nil {
				disputeLog.Errorw("disputer shutting down", "err", err)
				return err
			}
		}

		return nil
	},
}

// for a given miner, index, and maxPostIndex, tries to dispute posts from 0...postsSnapshotted-1
// returns a list of DisputeWindowedPoSt msgs that are expected to succeed if sent
func makeDisputeWindowedPosts(ctx context.Context, api v0api.FullNode, dl minerDeadline, postsSnapshotted uint64, sender address.Address) ([]*types.Message, error) {
	disputes := make([]*types.Message, 0)

	for i := uint64(0); i < postsSnapshotted; i++ {

		dpp, aerr := actors.SerializeParams(&miner3.DisputeWindowedPoStParams{
			Deadline:  dl.index,
			PoStIndex: i,
		})

		if aerr != nil {
			return nil, xerrors.Errorf("failed to serailize params: %w", aerr)
		}

		dispute := &types.Message{
			To:     dl.miner,
			From:   sender,
			Value:  big.Zero(),
			Method: builtin3.MethodsMiner.DisputeWindowedPoSt,
			Params: dpp,
		}

		rslt, err := api.StateCall(ctx, dispute, types.EmptyTSK)
		if err == nil && rslt.MsgRct.ExitCode == 0 {
			disputes = append(disputes, dispute)
		}

	}

	return disputes, nil
}

func makeMinerDeadline(ctx context.Context, api v0api.FullNode, mAddr address.Address) (abi.ChainEpoch, *minerDeadline, error) {
	dl, err := api.StateMinerProvingDeadline(ctx, mAddr, types.EmptyTSK)
	if err != nil {
		return -1, nil, xerrors.Errorf("getting proving index list: %w", err)
	}

	return dl.Close, &minerDeadline{
		miner: mAddr,
		index: dl.Index,
	}, nil
}

func getSender(ctx context.Context, api v0api.FullNode, fromStr string) (address.Address, error) {
	if fromStr == "" {
		return api.WalletDefaultAddress(ctx)
	}

	addr, err := address.NewFromString(fromStr)
	if err != nil {
		return address.Undef, err
	}

	has, err := api.WalletHas(ctx, addr)
	if err != nil {
		return address.Undef, err
	}

	if !has {
		return address.Undef, xerrors.Errorf("wallet doesn't contain: %s ", addr)
	}

	return addr, nil
}

func getMaxFee(maxStr string) (*lapi.MessageSendSpec, error) {
	if maxStr != "" {
		maxFee, err := types.ParseFIL(maxStr)
		if err != nil {
			return nil, xerrors.Errorf("parsing max-fee: %w", err)
		}
		return &lapi.MessageSendSpec{
			MaxFee: types.BigInt(maxFee),
		}, nil
	}

	return nil, nil
}
