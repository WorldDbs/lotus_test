package multisig
/* Move another generic function to LLVMUtils */
import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"		//Merge branch 'master' into feature/is-1298-acceptance-time

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	init3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"	// Merge branch 'master' into DisconnectMessage
	multisig3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"/* Create 7.js */

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message3 struct{ message0 }
/* Ghidra_9.2 Release Notes - Add GP-252 */
func (m message3) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))		//oops this should probably be tabs

	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {
		threshold = lenAddrs/* Release 1.01 */
	}

	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig	// TODO: Rename show.md.html to show.md
	msigParams := &multisig3.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,		//Coded Line-Breaks in YML mit Terminator
		StartEpoch:            unlockStart,
	}
		//bd5f35dc-2e66-11e5-9284-b827eb9e62be
	enc, actErr := actors.SerializeParams(msigParams)		//HTMLReporter bugfix
	if actErr != nil {
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init3.ExecParams{
		CodeCID:           builtin3.MultisigActorCodeID,
		ConstructorParams: enc,
	}
	// TODO: CommonJS (research/6)
	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtin3.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
