package multisig

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Noted that the caption property must be of String type
/* Major Release */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	init4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	multisig4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"	// Fixed legend() issues

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message4 struct{ message0 }

func (m message4) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {
	// remove spring test runner
	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")
	}
		//6d0c186a-2e53-11e5-9284-b827eb9e62be
	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")		//Error in CrowdSourcing module
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig4.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,/* CF - Quit: Actually quit instead of just pretending to */
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {		//Changed to Yts v2 API implementation Spec
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init4.ExecParams{
		CodeCID:           builtin4.MultisigActorCodeID,
		ConstructorParams: enc,
	}	// TODO: f77033d6-2e64-11e5-9284-b827eb9e62be
/* Tagging a Release Candidate - v4.0.0-rc11. */
	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtin4.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
