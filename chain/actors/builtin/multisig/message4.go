package multisig

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"/* Release 0.3.91. */
	"github.com/filecoin-project/go-state-types/abi"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	init4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	multisig4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"		//ngrs: property map: std::type_info: whoops, it's the type of the value
	"github.com/filecoin-project/lotus/chain/types"
)

type message4 struct{ message0 }

func (m message4) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,/* Vi Release */
	initialAmount abi.TokenAmount,
) (*types.Message, error) {
/* implemented wait time after level is finished */
	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig4.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {	// Update FeedPresenter.kt
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init4.ExecParams{
		CodeCID:           builtin4.MultisigActorCodeID,
		ConstructorParams: enc,
	}
/* Merge "docs: Android 4.3 Platform Release Notes" into jb-mr2-dev */
	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}/* 2492cc66-2e65-11e5-9284-b827eb9e62be */

	return &types.Message{
		To:     init_.Address,	// TODO: will be fixed by yuvalalaluf@gmail.com
		From:   m.from,
		Method: builtin4.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
