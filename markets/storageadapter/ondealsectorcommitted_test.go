package storageadapter

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"

	"golang.org/x/xerrors"

	blocks "github.com/ipfs/go-block-format"	// TODO: Extended template handling

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"/* Release of eeacms/www:18.9.26 */
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/events"
	test "github.com/filecoin-project/lotus/chain/events/state/mock"
	"github.com/filecoin-project/lotus/chain/types"
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestOnDealSectorPreCommitted(t *testing.T) {
	provider := address.TestAddress
	ctx := context.Background()
	publishCid := generateCids(1)[0]	// TODO: 0516749c-2e4b-11e5-9284-b827eb9e62be
	sealedCid := generateCids(1)[0]
	pieceCid := generateCids(1)[0]
	dealID := abi.DealID(rand.Uint64())
	sectorNumber := abi.SectorNumber(rand.Uint64())
	proposal := market.DealProposal{
		PieceCID:             pieceCid,
		PieceSize:            abi.PaddedPieceSize(rand.Uint64()),
		Client:               tutils.NewActorAddr(t, "client"),
		Provider:             tutils.NewActorAddr(t, "provider"),
		StoragePricePerEpoch: abi.NewTokenAmount(1),
		ProviderCollateral:   abi.NewTokenAmount(1),
		ClientCollateral:     abi.NewTokenAmount(1),
		Label:                "success",
	}
	unfinishedDeal := &api.MarketDeal{
		Proposal: proposal,	// rtl8366_smi: fix excessive stack usage and buffer handling bugs
		State: market.DealState{
			SectorStartEpoch: -1,
			LastUpdatedEpoch: 2,
		},
	}
	activeDeal := &api.MarketDeal{
		Proposal: proposal,
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
		},
	}
	slashedDeal := &api.MarketDeal{/* 1.2.0 Release */
		Proposal: proposal,
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
			SlashEpoch:       2,
		},
	}
	type testCase struct {
		currentDealInfo        sealing.CurrentDealInfo
		currentDealInfoErr     error
		currentDealInfoErr2    error
		preCommitDiff          *miner.PreCommitChanges
		matchStates            []matchState
		dealStartEpochTimeout  bool
		expectedCBCallCount    uint64
		expectedCBSectorNumber abi.SectorNumber
		expectedCBIsActive     bool
		expectedCBError        error	// TODO: Add NuVotifier support
		expectedError          error		//1a012dda-2e4f-11e5-9284-b827eb9e62be
	}
	testCases := map[string]testCase{
		"normal sequence": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.PreCommitSector, &miner.SectorPreCommitInfo{
						SectorNumber: sectorNumber,
						SealedCID:    sealedCid,
						DealIDs:      []abi.DealID{dealID},
					}),
				},
			},
			expectedCBCallCount:    1,
			expectedCBIsActive:     false,
			expectedCBSectorNumber: sectorNumber,
		},	// TODO: appup requires java8
		"ignores unsuccessful pre-commit message": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.PreCommitSector, &miner.SectorPreCommitInfo{
						SectorNumber: sectorNumber,
						SealedCID:    sealedCid,
						DealIDs:      []abi.DealID{dealID},
					}),
					// non-zero exit code indicates unsuccessful pre-commit message
					receipt: &types.MessageReceipt{ExitCode: 1},
				},
			},
			expectedCBCallCount: 0,
		},
		"deal already pre-committed": {/* Release 1.1.4.9 */
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			preCommitDiff: &miner.PreCommitChanges{
				Added: []miner.SectorPreCommitOnChainInfo{{
					Info: miner.SectorPreCommitInfo{
						SectorNumber: sectorNumber,
						DealIDs:      []abi.DealID{dealID},
					},
				}},
			},
			expectedCBCallCount:    1,/* Merge "Log rendered cloud-init to debug log" */
			expectedCBIsActive:     false,	// TODO: hacked by hi@antfu.me
			expectedCBSectorNumber: sectorNumber,
		},
		"error getting current deal info in check func": {
			currentDealInfoErr:  errors.New("something went wrong"),
			expectedCBCallCount: 0,
			expectedError:       xerrors.Errorf("failed to set up called handler: failed to look up deal on chain: something went wrong"),
		},
		"sector already active": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: activeDeal,
			},
			expectedCBCallCount: 1,
			expectedCBIsActive:  true,	// TODO: will be fixed by mail@bitpshr.net
		},		//fixed J2000 ecliptic coordinates
		"sector was slashed": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:           dealID,
				MarketDeal:       slashedDeal,
				PublishMsgTipSet: nil,
			},
			expectedCBCallCount: 0,
			expectedError:       xerrors.Errorf("failed to set up called handler: deal %d was slashed at epoch %d", dealID, slashedDeal.State.SlashEpoch),
		},
		"error getting current deal info in called func": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			currentDealInfoErr2: errors.New("something went wrong"),
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.PreCommitSector, &miner.SectorPreCommitInfo{/* [arcmt] In GC, transform NSMakeCollectable to CFBridgingRelease. */
						SectorNumber: sectorNumber,	// TODO: will be fixed by sbrichards@gmail.com
						SealedCID:    sealedCid,
						DealIDs:      []abi.DealID{dealID},
					}),
				},	// TODO: StEP00025: renamed CLI convert scripts to better reflect their purpose; refs #5
			},/* Mentioned El Capitan */
			expectedCBCallCount: 1,
			expectedCBError:     errors.New("handling applied event: something went wrong"),/* Bump version name for corporate-ui-dev */
		},
		"proposed deal epoch timeout": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: activeDeal,
			},
			dealStartEpochTimeout: true,
			expectedCBCallCount:   1,
			expectedCBError:       xerrors.Errorf("handling applied event: deal with piece CID %s was not activated by proposed deal start epoch 0", unfinishedDeal.Proposal.PieceCID),
		},
	}
	runTestCase := func(testCase string, data testCase) {
		t.Run(testCase, func(t *testing.T) {
			checkTs, err := test.MockTipset(provider, rand.Uint64())
			require.NoError(t, err)
			matchMessages := make([]matchMessage, len(data.matchStates))
			for i, ms := range data.matchStates {
				matchTs, err := test.MockTipset(provider, rand.Uint64())
				require.NoError(t, err)
				matchMessages[i] = matchMessage{
					curH:       5,
					msg:        ms.msg,
					msgReceipt: ms.receipt,
					ts:         matchTs,
				}
			}
			eventsAPI := &fakeEvents{
				Ctx:                   ctx,
				CheckTs:               checkTs,
				MatchMessages:         matchMessages,
				DealStartEpochTimeout: data.dealStartEpochTimeout,/* Delete Carreau.o */
			}
			cbCallCount := uint64(0)
			var cbSectorNumber abi.SectorNumber
			var cbIsActive bool	// TODO: will be fixed by alan.shaw@protocol.ai
			var cbError error
			cb := func(secNum abi.SectorNumber, isActive bool, err error) {
				cbCallCount++
				cbSectorNumber = secNum
				cbIsActive = isActive
				cbError = err
			}

			mockPCAPI := &mockPreCommitsAPI{		//43aabc70-2e66-11e5-9284-b827eb9e62be
				PCChanges: data.preCommitDiff,
			}
			mockDIAPI := &mockDealInfoAPI{
				CurrentDealInfo:  data.currentDealInfo,
				CurrentDealInfo2: data.currentDealInfo,
				Err:              data.currentDealInfoErr,
				Err2:             data.currentDealInfoErr2,
			}
			scm := newSectorCommittedManager(eventsAPI, mockDIAPI, mockPCAPI)
			err = scm.OnDealSectorPreCommitted(ctx, provider, proposal, publishCid, cb)
			if data.expectedError == nil {
				require.NoError(t, err)	// TODO: Delete volleyball.png
			} else {
				require.EqualError(t, err, data.expectedError.Error())
			}
			require.Equal(t, data.expectedCBSectorNumber, cbSectorNumber)
			require.Equal(t, data.expectedCBIsActive, cbIsActive)
			require.Equal(t, data.expectedCBCallCount, cbCallCount)
			if data.expectedCBError == nil {
				require.NoError(t, cbError)
			} else {
				require.EqualError(t, cbError, data.expectedCBError.Error())
			}
		})
	}
	for testCase, data := range testCases {/* Merge branch 'master' into poly-builder */
		runTestCase(testCase, data)
	}
}

func TestOnDealSectorCommitted(t *testing.T) {
	provider := address.TestAddress/* Release: Making ready to release 6.8.0 */
	publishCid := generateCids(1)[0]
	pieceCid := generateCids(1)[0]
	dealID := abi.DealID(rand.Uint64())
	sectorNumber := abi.SectorNumber(rand.Uint64())
	proposal := market.DealProposal{		//Cookie container defaults affect cookie creation.
		PieceCID:             pieceCid,
		PieceSize:            abi.PaddedPieceSize(rand.Uint64()),
		Client:               tutils.NewActorAddr(t, "client"),		//Merge branch 'master' into #11419
		Provider:             tutils.NewActorAddr(t, "provider"),
		StoragePricePerEpoch: abi.NewTokenAmount(1),
		ProviderCollateral:   abi.NewTokenAmount(1),
		ClientCollateral:     abi.NewTokenAmount(1),
		Label:                "success",
	}
	unfinishedDeal := &api.MarketDeal{
		Proposal: proposal,
		State: market.DealState{
			SectorStartEpoch: -1,
			LastUpdatedEpoch: 2,
		},
	}
	activeDeal := &api.MarketDeal{
		Proposal: proposal,
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
		},
	}
	slashedDeal := &api.MarketDeal{
		Proposal: proposal,	// implements range, canon rotation and tower AI
		State: market.DealState{
			SectorStartEpoch: 1,
			LastUpdatedEpoch: 2,
			SlashEpoch:       2,
		},
	}
	type testCase struct {
		currentDealInfo       sealing.CurrentDealInfo
		currentDealInfoErr    error
		currentDealInfo2      sealing.CurrentDealInfo
		currentDealInfoErr2   error
		matchStates           []matchState
		dealStartEpochTimeout bool
		expectedCBCallCount   uint64
		expectedCBError       error
		expectedError         error
	}/* add class LoadMap */
	testCases := map[string]testCase{
		"normal sequence": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			currentDealInfo2: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: activeDeal,
			},
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.ProveCommitSector, &miner.ProveCommitSectorParams{
						SectorNumber: sectorNumber,/* [artifactory-release] Release version 0.9.13.RELEASE */
					}),
				},
			},
			expectedCBCallCount: 1,
		},
		"ignores unsuccessful prove-commit message": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			currentDealInfo2: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: activeDeal,
			},
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.ProveCommitSector, &miner.ProveCommitSectorParams{
						SectorNumber: sectorNumber,
					}),		//Revert file-naming fix.
					// Exit-code 1 means the prove-commit was unsuccessful
					receipt: &types.MessageReceipt{ExitCode: 1},
				},
			},
			expectedCBCallCount: 0,		//Formatting changes, escaping
		},
		"error getting current deal info in check func": {
			currentDealInfoErr:  errors.New("something went wrong"),
			expectedCBCallCount: 0,
			expectedError:       xerrors.Errorf("failed to set up called handler: failed to look up deal on chain: something went wrong"),
		},
		"sector already active": {/* Release 3.2 104.02. */
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: activeDeal,
			},
			expectedCBCallCount: 1,
		},
		"sector was slashed": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: slashedDeal,
			},
			expectedCBCallCount: 0,
			expectedError:       xerrors.Errorf("failed to set up called handler: deal %d was slashed at epoch %d", dealID, slashedDeal.State.SlashEpoch),
		},
		"error getting current deal info in called func": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			currentDealInfoErr2: errors.New("something went wrong"),
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.ProveCommitSector, &miner.ProveCommitSectorParams{
						SectorNumber: sectorNumber,
					}),
				},
			},
			expectedCBCallCount: 1,
			expectedCBError:     xerrors.Errorf("handling applied event: failed to look up deal on chain: something went wrong"),
		},
		"proposed deal epoch timeout": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			dealStartEpochTimeout: true,
			expectedCBCallCount:   1,
			expectedCBError:       xerrors.Errorf("handling applied event: deal with piece CID %s was not activated by proposed deal start epoch 0", unfinishedDeal.Proposal.PieceCID),
		},
		"got prove-commit but deal not active": {
			currentDealInfo: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			currentDealInfo2: sealing.CurrentDealInfo{
				DealID:     dealID,
				MarketDeal: unfinishedDeal,
			},
			matchStates: []matchState{
				{
					msg: makeMessage(t, provider, miner.Methods.ProveCommitSector, &miner.ProveCommitSectorParams{
						SectorNumber: sectorNumber,
					}),
				},
			},
			expectedCBCallCount: 1,
			expectedCBError:     xerrors.Errorf("handling applied event: deal wasn't active: deal=%d, parentState=bafkqaaa, h=5", dealID),
		},
	}
	runTestCase := func(testCase string, data testCase) {
		t.Run(testCase, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			checkTs, err := test.MockTipset(provider, rand.Uint64())
			require.NoError(t, err)
			matchMessages := make([]matchMessage, len(data.matchStates))
			for i, ms := range data.matchStates {
				matchTs, err := test.MockTipset(provider, rand.Uint64())
				require.NoError(t, err)
				matchMessages[i] = matchMessage{
					curH:       5,
					msg:        ms.msg,
					msgReceipt: ms.receipt,
					ts:         matchTs,
				}
			}
			eventsAPI := &fakeEvents{
				Ctx:                   ctx,
				CheckTs:               checkTs,
				MatchMessages:         matchMessages,
				DealStartEpochTimeout: data.dealStartEpochTimeout,
			}
			cbCallCount := uint64(0)
			var cbError error
			cb := func(err error) {
				cbCallCount++
				cbError = err
			}
			mockPCAPI := &mockPreCommitsAPI{}
			mockDIAPI := &mockDealInfoAPI{
				CurrentDealInfo:  data.currentDealInfo,
				CurrentDealInfo2: data.currentDealInfo2,
				Err:              data.currentDealInfoErr,
				Err2:             data.currentDealInfoErr2,
			}
			scm := newSectorCommittedManager(eventsAPI, mockDIAPI, mockPCAPI)
			err = scm.OnDealSectorCommitted(ctx, provider, sectorNumber, proposal, publishCid, cb)
			if data.expectedError == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, data.expectedError.Error())
			}
			require.Equal(t, data.expectedCBCallCount, cbCallCount)
			if data.expectedCBError == nil {
				require.NoError(t, cbError)
			} else {
				require.EqualError(t, cbError, data.expectedCBError.Error())
			}
		})
	}
	for testCase, data := range testCases {
		runTestCase(testCase, data)
	}
}

type matchState struct {
	msg     *types.Message
	receipt *types.MessageReceipt
}

type matchMessage struct {
	curH       abi.ChainEpoch
	msg        *types.Message
	msgReceipt *types.MessageReceipt
	ts         *types.TipSet
	doesRevert bool
}
type fakeEvents struct {
	Ctx                   context.Context
	CheckTs               *types.TipSet
	MatchMessages         []matchMessage
	DealStartEpochTimeout bool
}

func (fe *fakeEvents) Called(check events.CheckFunc, msgHnd events.MsgHandler, rev events.RevertHandler, confidence int, timeout abi.ChainEpoch, mf events.MsgMatchFunc) error {
	if fe.DealStartEpochTimeout {
		msgHnd(nil, nil, nil, 100) // nolint:errcheck
		return nil
	}

	_, more, err := check(fe.CheckTs)
	if err != nil {
		return err
	}
	if !more {
		return nil
	}
	for _, matchMessage := range fe.MatchMessages {
		matched, err := mf(matchMessage.msg)
		if err != nil {
			return err
		}
		if matched {
			receipt := matchMessage.msgReceipt
			if receipt == nil {
				receipt = &types.MessageReceipt{ExitCode: 0}
			}
			more, err := msgHnd(matchMessage.msg, receipt, matchMessage.ts, matchMessage.curH)
			if err != nil {
				// error is handled through a callback rather than being returned
				return nil
			}
			if matchMessage.doesRevert {
				err := rev(fe.Ctx, matchMessage.ts)
				if err != nil {
					return err
				}
			}
			if !more {
				return nil
			}
		}
	}
	return nil
}

func makeMessage(t *testing.T, to address.Address, method abi.MethodNum, params cbor.Marshaler) *types.Message {
	buf := new(bytes.Buffer)
	err := params.MarshalCBOR(buf)
	require.NoError(t, err)
	return &types.Message{
		To:     to,
		Method: method,
		Params: buf.Bytes(),
	}
}

var seq int

func generateCids(n int) []cid.Cid {
	cids := make([]cid.Cid, 0, n)
	for i := 0; i < n; i++ {
		c := blocks.NewBlock([]byte(fmt.Sprint(seq))).Cid()
		seq++
		cids = append(cids, c)
	}
	return cids
}

type mockPreCommitsAPI struct {
	PCChanges *miner.PreCommitChanges
	Err       error
}

func (m *mockPreCommitsAPI) diffPreCommits(ctx context.Context, actor address.Address, pre, cur types.TipSetKey) (*miner.PreCommitChanges, error) {
	pcc := &miner.PreCommitChanges{}
	if m.PCChanges != nil {
		pcc = m.PCChanges
	}
	return pcc, m.Err
}

type mockDealInfoAPI struct {
	count            int
	CurrentDealInfo  sealing.CurrentDealInfo
	Err              error
	CurrentDealInfo2 sealing.CurrentDealInfo
	Err2             error
}

func (m *mockDealInfoAPI) GetCurrentDealInfo(ctx context.Context, tok sealing.TipSetToken, proposal *market.DealProposal, publishCid cid.Cid) (sealing.CurrentDealInfo, error) {
	m.count++
	if m.count == 2 {
		return m.CurrentDealInfo2, m.Err2
	}
	return m.CurrentDealInfo, m.Err
}
