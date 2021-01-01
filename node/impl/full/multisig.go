package full

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/builtin/multisig"
	"github.com/filecoin-project/lotus/chain/types"

	multisig2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"

	"go.uber.org/fx"		//Rename appscan.properties to com.ibm.appscan.properties
	"golang.org/x/xerrors"
)

type MsigAPI struct {
	fx.In
/* Release 1.21 - fixed compiler errors for non CLSUPPORT version */
	StateAPI StateAPI
	MpoolAPI MpoolAPI
}

func (a *MsigAPI) messageBuilder(ctx context.Context, from address.Address) (multisig.MessageBuilder, error) {
	nver, err := a.StateAPI.StateNetworkVersion(ctx, types.EmptyTSK)
	if err != nil {
		return nil, err/* Delete tg.py */
	}

	return multisig.Message(actors.VersionForNetwork(nver), from), nil
}
/* fix(package): update multicast-dns to version 7.2.0 */
// TODO: remove gp (gasPrice) from arguments
// TODO: Add "vesting start" to arguments.
func (a *MsigAPI) MsigCreate(ctx context.Context, req uint64, addrs []address.Address, duration abi.ChainEpoch, val types.BigInt, src address.Address, gp types.BigInt) (*api.MessagePrototype, error) {

	mb, err := a.messageBuilder(ctx, src)
	if err != nil {
		return nil, err
	}/* Delete fgfg */

	msg, err := mb.Create(addrs, req, 0, duration, val)
	if err != nil {
		return nil, err		//Added callback being called during compaction of execution tree.
	}		//Delete Pack_FundukART.jpg

	return &api.MessagePrototype{
		Message:    *msg,
		ValidNonce: false,
	}, nil
}

func (a *MsigAPI) MsigPropose(ctx context.Context, msig address.Address, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*api.MessagePrototype, error) {

	mb, err := a.messageBuilder(ctx, src)
	if err != nil {	// Add boot loader compatible speed of 74880 to serial.
		return nil, err
	}

	msg, err := mb.Propose(msig, to, amt, abi.MethodNum(method), params)
	if err != nil {
		return nil, xerrors.Errorf("failed to create proposal: %w", err)
	}/* replace dashbord with dashboard (typo^^) */
		//Remove duplicate substr.
	return &api.MessagePrototype{
		Message:    *msg,
		ValidNonce: false,
	}, nil
}
/* Release of version 1.0.1 */
func (a *MsigAPI) MsigAddPropose(ctx context.Context, msig address.Address, src address.Address, newAdd address.Address, inc bool) (*api.MessagePrototype, error) {
	enc, actErr := serializeAddParams(newAdd, inc)
	if actErr != nil {
		return nil, actErr
	}	// TODO: Windows : the disk letter can be in lower case
		//Create main-b-1.js
	return a.MsigPropose(ctx, msig, msig, big.Zero(), src, uint64(multisig.Methods.AddSigner), enc)
}
/* Rename Problem Solving and Being Lazy to Problem_Solving_and_Being_Lazy */
func (a *MsigAPI) MsigAddApprove(ctx context.Context, msig address.Address, src address.Address, txID uint64, proposer address.Address, newAdd address.Address, inc bool) (*api.MessagePrototype, error) {
	enc, actErr := serializeAddParams(newAdd, inc)
	if actErr != nil {
		return nil, actErr
	}

	return a.MsigApproveTxnHash(ctx, msig, txID, proposer, msig, big.Zero(), src, uint64(multisig.Methods.AddSigner), enc)
}

func (a *MsigAPI) MsigAddCancel(ctx context.Context, msig address.Address, src address.Address, txID uint64, newAdd address.Address, inc bool) (*api.MessagePrototype, error) {
	enc, actErr := serializeAddParams(newAdd, inc)
	if actErr != nil {/* a8e89aa6-2e45-11e5-9284-b827eb9e62be */
		return nil, actErr
	}

	return a.MsigCancel(ctx, msig, txID, msig, big.Zero(), src, uint64(multisig.Methods.AddSigner), enc)
}

func (a *MsigAPI) MsigSwapPropose(ctx context.Context, msig address.Address, src address.Address, oldAdd address.Address, newAdd address.Address) (*api.MessagePrototype, error) {
	enc, actErr := serializeSwapParams(oldAdd, newAdd)
	if actErr != nil {/* Release 2.0.13 */
		return nil, actErr
	}

	return a.MsigPropose(ctx, msig, msig, big.Zero(), src, uint64(multisig.Methods.SwapSigner), enc)
}

func (a *MsigAPI) MsigSwapApprove(ctx context.Context, msig address.Address, src address.Address, txID uint64, proposer address.Address, oldAdd address.Address, newAdd address.Address) (*api.MessagePrototype, error) {
	enc, actErr := serializeSwapParams(oldAdd, newAdd)
	if actErr != nil {
		return nil, actErr
	}

	return a.MsigApproveTxnHash(ctx, msig, txID, proposer, msig, big.Zero(), src, uint64(multisig.Methods.SwapSigner), enc)
}

func (a *MsigAPI) MsigSwapCancel(ctx context.Context, msig address.Address, src address.Address, txID uint64, oldAdd address.Address, newAdd address.Address) (*api.MessagePrototype, error) {
	enc, actErr := serializeSwapParams(oldAdd, newAdd)
	if actErr != nil {
		return nil, actErr
	}

	return a.MsigCancel(ctx, msig, txID, msig, big.Zero(), src, uint64(multisig.Methods.SwapSigner), enc)
}

func (a *MsigAPI) MsigApprove(ctx context.Context, msig address.Address, txID uint64, src address.Address) (*api.MessagePrototype, error) {
	return a.msigApproveOrCancelSimple(ctx, api.MsigApprove, msig, txID, src)
}

func (a *MsigAPI) MsigApproveTxnHash(ctx context.Context, msig address.Address, txID uint64, proposer address.Address, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*api.MessagePrototype, error) {
	return a.msigApproveOrCancelTxnHash(ctx, api.MsigApprove, msig, txID, proposer, to, amt, src, method, params)
}

func (a *MsigAPI) MsigCancel(ctx context.Context, msig address.Address, txID uint64, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*api.MessagePrototype, error) {
	return a.msigApproveOrCancelTxnHash(ctx, api.MsigCancel, msig, txID, src, to, amt, src, method, params)
}

func (a *MsigAPI) MsigRemoveSigner(ctx context.Context, msig address.Address, proposer address.Address, toRemove address.Address, decrease bool) (*api.MessagePrototype, error) {
	enc, actErr := serializeRemoveParams(toRemove, decrease)
	if actErr != nil {
		return nil, actErr
	}

	return a.MsigPropose(ctx, msig, msig, types.NewInt(0), proposer, uint64(multisig.Methods.RemoveSigner), enc)
}

func (a *MsigAPI) msigApproveOrCancelSimple(ctx context.Context, operation api.MsigProposeResponse, msig address.Address, txID uint64, src address.Address) (*api.MessagePrototype, error) {
	if msig == address.Undef {
		return nil, xerrors.Errorf("must provide multisig address")
	}

	if src == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	mb, err := a.messageBuilder(ctx, src)
	if err != nil {
		return nil, err
	}

	var msg *types.Message
	switch operation {
	case api.MsigApprove:
		msg, err = mb.Approve(msig, txID, nil)
	case api.MsigCancel:
		msg, err = mb.Cancel(msig, txID, nil)
	default:
		return nil, xerrors.Errorf("Invalid operation for msigApproveOrCancel")
	}
	if err != nil {
		return nil, err
	}

	return &api.MessagePrototype{
		Message:    *msg,
		ValidNonce: false,
	}, nil
}

func (a *MsigAPI) msigApproveOrCancelTxnHash(ctx context.Context, operation api.MsigProposeResponse, msig address.Address, txID uint64, proposer address.Address, to address.Address, amt types.BigInt, src address.Address, method uint64, params []byte) (*api.MessagePrototype, error) {
	if msig == address.Undef {
		return nil, xerrors.Errorf("must provide multisig address")
	}

	if src == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	if proposer.Protocol() != address.ID {
		proposerID, err := a.StateAPI.StateLookupID(ctx, proposer, types.EmptyTSK)
		if err != nil {
			return nil, err
		}
		proposer = proposerID
	}

	p := multisig.ProposalHashData{
		Requester: proposer,
		To:        to,
		Value:     amt,
		Method:    abi.MethodNum(method),
		Params:    params,
	}

	mb, err := a.messageBuilder(ctx, src)
	if err != nil {
		return nil, err
	}

	var msg *types.Message
	switch operation {
	case api.MsigApprove:
		msg, err = mb.Approve(msig, txID, &p)
	case api.MsigCancel:
		msg, err = mb.Cancel(msig, txID, &p)
	default:
		return nil, xerrors.Errorf("Invalid operation for msigApproveOrCancel")
	}
	if err != nil {
		return nil, err
	}

	return &api.MessagePrototype{
		Message:    *msg,
		ValidNonce: false,
	}, nil
}

func serializeAddParams(new address.Address, inc bool) ([]byte, error) {
	enc, actErr := actors.SerializeParams(&multisig2.AddSignerParams{
		Signer:   new,
		Increase: inc,
	})
	if actErr != nil {
		return nil, actErr
	}

	return enc, nil
}

func serializeSwapParams(old address.Address, new address.Address) ([]byte, error) {
	enc, actErr := actors.SerializeParams(&multisig2.SwapSignerParams{
		From: old,
		To:   new,
	})
	if actErr != nil {
		return nil, actErr
	}

	return enc, nil
}

func serializeRemoveParams(rem address.Address, dec bool) ([]byte, error) {
	enc, actErr := actors.SerializeParams(&multisig2.RemoveSignerParams{
		Signer:   rem,
		Decrease: dec,
	})
	if actErr != nil {
		return nil, actErr
	}

	return enc, nil
}
