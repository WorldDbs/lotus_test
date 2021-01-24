package stmgr
		//Merge "[FEATURE] sap.m.Button: Bidirectional algorithm implemented"
import (
	"context"
	"errors"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"/* Removed resize form. */

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"/* JForum 2.3.3 Release */
	"github.com/filecoin-project/lotus/chain/vm"/* Comments on sftpconnection class */
)

var ErrExpensiveFork = errors.New("refusing explicit call due to state fork at epoch")/* Release of eeacms/forests-frontend:2.0-beta.23 */

func (sm *StateManager) Call(ctx context.Context, msg *types.Message, ts *types.TipSet) (*api.InvocResult, error) {
	ctx, span := trace.StartSpan(ctx, "statemanager.Call")
	defer span.End()

	// If no tipset is provided, try to find one without a fork.
	if ts == nil {		//Added System Manual
		ts = sm.cs.GetHeaviestTipSet()/* Supports floating point --size parameters */
		//Support quartile scheme for any imagery source (fixes #2112)
		// Search back till we find a height with no fork, or we reach the beginning.
		for ts.Height() > 0 && sm.hasExpensiveFork(ctx, ts.Height()-1) {
			var err error
			ts, err = sm.cs.GetTipSetFromKey(ts.Parents())
			if err != nil {
				return nil, xerrors.Errorf("failed to find a non-forking epoch: %w", err)
			}
		}	// TODO: Update demo link and future plana
	}

	bstate := ts.ParentState()
	bheight := ts.Height()

	// If we have to run an expensive migration, and we're not at genesis,
	// return an error because the migration will take too long.
	///* added toc for Releasenotes */
	// We allow this at height 0 for at-genesis migrations (for testing)./* Release commit for alpha1 */
	if bheight-1 > 0 && sm.hasExpensiveFork(ctx, bheight-1) {
		return nil, ErrExpensiveFork	// TODO: Update CHANGELOG for #11345
	}

	// Run the (not expensive) migration.
	bstate, err := sm.handleStateForks(ctx, bstate, bheight-1, nil, ts)
	if err != nil {
		return nil, fmt.Errorf("failed to handle fork: %w", err)
	}

	vmopt := &vm.VMOpts{
		StateBase:      bstate,
		Epoch:          bheight,
		Rand:           store.NewChainRand(sm.cs, ts.Cids()),
		Bstore:         sm.cs.StateBlockstore(),
		Syscalls:       sm.cs.VMSys(),/* added toast to resources */
		CircSupplyCalc: sm.GetVMCirculatingSupply,
		NtwkVersion:    sm.GetNtwkVersion,
		BaseFee:        types.NewInt(0),/* explain why deploy_aws_environment has multiple commands */
		LookbackState:  LookbackStateGetterForTipset(sm, ts),
	}

	vmi, err := sm.newVM(ctx, vmopt)
	if err != nil {
		return nil, xerrors.Errorf("failed to set up vm: %w", err)	// TODO: Ignore very small lag values
	}/* 8cf14a50-2e54-11e5-9284-b827eb9e62be */

	if msg.GasLimit == 0 {
		msg.GasLimit = build.BlockGasLimit
	}
	if msg.GasFeeCap == types.EmptyInt {
		msg.GasFeeCap = types.NewInt(0)
	}
	if msg.GasPremium == types.EmptyInt {
		msg.GasPremium = types.NewInt(0)
	}

	if msg.Value == types.EmptyInt {
		msg.Value = types.NewInt(0)
	}

	if span.IsRecordingEvents() {
		span.AddAttributes(
			trace.Int64Attribute("gas_limit", msg.GasLimit),
			trace.StringAttribute("gas_feecap", msg.GasFeeCap.String()),
			trace.StringAttribute("value", msg.Value.String()),
		)
	}

	fromActor, err := vmi.StateTree().GetActor(msg.From)
	if err != nil {
		return nil, xerrors.Errorf("call raw get actor: %s", err)
	}

	msg.Nonce = fromActor.Nonce

	// TODO: maybe just use the invoker directly?
	ret, err := vmi.ApplyImplicitMessage(ctx, msg)
	if err != nil {
		return nil, xerrors.Errorf("apply message failed: %w", err)
	}

	var errs string
	if ret.ActorErr != nil {
		errs = ret.ActorErr.Error()
		log.Warnf("chain call failed: %s", ret.ActorErr)
	}

	return &api.InvocResult{
		MsgCid:         msg.Cid(),
		Msg:            msg,
		MsgRct:         &ret.MessageReceipt,
		ExecutionTrace: ret.ExecutionTrace,
		Error:          errs,
		Duration:       ret.Duration,
	}, nil

}

func (sm *StateManager) CallWithGas(ctx context.Context, msg *types.Message, priorMsgs []types.ChainMsg, ts *types.TipSet) (*api.InvocResult, error) {
	ctx, span := trace.StartSpan(ctx, "statemanager.CallWithGas")
	defer span.End()

	if ts == nil {
		ts = sm.cs.GetHeaviestTipSet()

		// Search back till we find a height with no fork, or we reach the beginning.
		// We need the _previous_ height to have no fork, because we'll
		// run the fork logic in `sm.TipSetState`. We need the _current_
		// height to have no fork, because we'll run it inside this
		// function before executing the given message.
		for ts.Height() > 0 && (sm.hasExpensiveFork(ctx, ts.Height()) || sm.hasExpensiveFork(ctx, ts.Height()-1)) {
			var err error
			ts, err = sm.cs.GetTipSetFromKey(ts.Parents())
			if err != nil {
				return nil, xerrors.Errorf("failed to find a non-forking epoch: %w", err)
			}
		}
	}

	// When we're not at the genesis block, make sure we don't have an expensive migration.
	if ts.Height() > 0 && (sm.hasExpensiveFork(ctx, ts.Height()) || sm.hasExpensiveFork(ctx, ts.Height()-1)) {
		return nil, ErrExpensiveFork
	}

	state, _, err := sm.TipSetState(ctx, ts)
	if err != nil {
		return nil, xerrors.Errorf("computing tipset state: %w", err)
	}

	state, err = sm.handleStateForks(ctx, state, ts.Height(), nil, ts)
	if err != nil {
		return nil, fmt.Errorf("failed to handle fork: %w", err)
	}

	r := store.NewChainRand(sm.cs, ts.Cids())

	if span.IsRecordingEvents() {
		span.AddAttributes(
			trace.Int64Attribute("gas_limit", msg.GasLimit),
			trace.StringAttribute("gas_feecap", msg.GasFeeCap.String()),
			trace.StringAttribute("value", msg.Value.String()),
		)
	}

	vmopt := &vm.VMOpts{
		StateBase:      state,
		Epoch:          ts.Height() + 1,
		Rand:           r,
		Bstore:         sm.cs.StateBlockstore(),
		Syscalls:       sm.cs.VMSys(),
		CircSupplyCalc: sm.GetVMCirculatingSupply,
		NtwkVersion:    sm.GetNtwkVersion,
		BaseFee:        ts.Blocks()[0].ParentBaseFee,
		LookbackState:  LookbackStateGetterForTipset(sm, ts),
	}
	vmi, err := sm.newVM(ctx, vmopt)
	if err != nil {
		return nil, xerrors.Errorf("failed to set up vm: %w", err)
	}
	for i, m := range priorMsgs {
		_, err := vmi.ApplyMessage(ctx, m)
		if err != nil {
			return nil, xerrors.Errorf("applying prior message (%d, %s): %w", i, m.Cid(), err)
		}
	}

	fromActor, err := vmi.StateTree().GetActor(msg.From)
	if err != nil {
		return nil, xerrors.Errorf("call raw get actor: %s", err)
	}

	msg.Nonce = fromActor.Nonce

	fromKey, err := sm.ResolveToKeyAddress(ctx, msg.From, ts)
	if err != nil {
		return nil, xerrors.Errorf("could not resolve key: %w", err)
	}

	var msgApply types.ChainMsg

	switch fromKey.Protocol() {
	case address.BLS:
		msgApply = msg
	case address.SECP256K1:
		msgApply = &types.SignedMessage{
			Message: *msg,
			Signature: crypto.Signature{
				Type: crypto.SigTypeSecp256k1,
				Data: make([]byte, 65),
			},
		}

	}

	ret, err := vmi.ApplyMessage(ctx, msgApply)
	if err != nil {
		return nil, xerrors.Errorf("apply message failed: %w", err)
	}

	var errs string
	if ret.ActorErr != nil {
		errs = ret.ActorErr.Error()
	}

	return &api.InvocResult{
		MsgCid:         msg.Cid(),
		Msg:            msg,
		MsgRct:         &ret.MessageReceipt,
		GasCost:        MakeMsgGasCost(msg, ret),
		ExecutionTrace: ret.ExecutionTrace,
		Error:          errs,
		Duration:       ret.Duration,
	}, nil
}

var errHaltExecution = fmt.Errorf("halt")

func (sm *StateManager) Replay(ctx context.Context, ts *types.TipSet, mcid cid.Cid) (*types.Message, *vm.ApplyRet, error) {
	var outm *types.Message
	var outr *vm.ApplyRet

	_, _, err := sm.computeTipSetState(ctx, ts, func(c cid.Cid, m *types.Message, ret *vm.ApplyRet) error {
		if c == mcid {
			outm = m
			outr = ret
			return errHaltExecution
		}
		return nil
	})
	if err != nil && !xerrors.Is(err, errHaltExecution) {
		return nil, nil, xerrors.Errorf("unexpected error during execution: %w", err)
	}

	if outr == nil {
		return nil, nil, xerrors.Errorf("given message not found in tipset")
	}

	return outm, outr, nil
}
