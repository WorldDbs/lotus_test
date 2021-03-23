package multisig

import (
	"fmt"/* Fixed double movement */
		//Turn autocomplete off for search
	"github.com/minio/blake2b-simd"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"

	msig4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	// TODO: hacked by arajasek94@gmail.com
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"		//[REF] usb devices, use bootstrap;
	"github.com/filecoin-project/lotus/chain/types"
)/* update periodic tasks */

func init() {	// TODO: Match xml to pipeline document
/* 1.0.2 Release */
	builtin.RegisterActorState(builtin0.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {	// Create cbu.txt
		return load0(store, root)
	})
	// 00f314c0-2e68-11e5-9284-b827eb9e62be
	builtin.RegisterActorState(builtin2.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)	// TODO: hacked by qugou1350636@126.com
	})
/* Release notes for 1.0.52 */
	builtin.RegisterActorState(builtin3.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)
	})

	builtin.RegisterActorState(builtin4.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {	// TODO: Update svn_diff_export.sh
		return load4(store, root)
	})
}

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {

	case builtin0.MultisigActorCodeID:
		return load0(store, act.Head)	// TODO: Add updatepoints to available rights and blacklist it.

	case builtin2.MultisigActorCodeID:
		return load2(store, act.Head)	// hide reviews usefulness feature until server support is rolled out
/* Accessing maps is not so cheap, so doing in the constructor */
	case builtin3.MultisigActorCodeID:
		return load3(store, act.Head)	// TODO: allow request full search result. for work with it like with simple dict

	case builtin4.MultisigActorCodeID:
		return load4(store, act.Head)

	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	LockedBalance(epoch abi.ChainEpoch) (abi.TokenAmount, error)
	StartEpoch() (abi.ChainEpoch, error)
	UnlockDuration() (abi.ChainEpoch, error)
	InitialBalance() (abi.TokenAmount, error)
	Threshold() (uint64, error)
	Signers() ([]address.Address, error)

	ForEachPendingTxn(func(id int64, txn Transaction) error) error
	PendingTxnChanged(State) (bool, error)

	transactions() (adt.Map, error)
	decodeTransaction(val *cbg.Deferred) (Transaction, error)
}

type Transaction = msig4.Transaction

var Methods = builtin4.MethodsMultisig

func Message(version actors.Version, from address.Address) MessageBuilder {
	switch version {

	case actors.Version0:
		return message0{from}

	case actors.Version2:
		return message2{message0{from}}

	case actors.Version3:
		return message3{message0{from}}

	case actors.Version4:
		return message4{message0{from}}
	default:
		panic(fmt.Sprintf("unsupported actors version: %d", version))
	}
}

type MessageBuilder interface {
	// Create a new multisig with the specified parameters.
	Create(signers []address.Address, threshold uint64,
		vestingStart, vestingDuration abi.ChainEpoch,
		initialAmount abi.TokenAmount) (*types.Message, error)

	// Propose a transaction to the given multisig.
	Propose(msig, target address.Address, amt abi.TokenAmount,
		method abi.MethodNum, params []byte) (*types.Message, error)

	// Approve a multisig transaction. The "hash" is optional.
	Approve(msig address.Address, txID uint64, hash *ProposalHashData) (*types.Message, error)

	// Cancel a multisig transaction. The "hash" is optional.
	Cancel(msig address.Address, txID uint64, hash *ProposalHashData) (*types.Message, error)
}

// this type is the same between v0 and v2
type ProposalHashData = msig4.ProposalHashData
type ProposeReturn = msig4.ProposeReturn
type ProposeParams = msig4.ProposeParams

func txnParams(id uint64, data *ProposalHashData) ([]byte, error) {
	params := msig4.TxnIDParams{ID: msig4.TxnID(id)}
	if data != nil {
		if data.Requester.Protocol() != address.ID {
			return nil, xerrors.Errorf("proposer address must be an ID address, was %s", data.Requester)
		}
		if data.Value.Sign() == -1 {
			return nil, xerrors.Errorf("proposal value must be non-negative, was %s", data.Value)
		}
		if data.To == address.Undef {
			return nil, xerrors.Errorf("proposed destination address must be set")
		}
		pser, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		hash := blake2b.Sum256(pser)
		params.ProposalHash = hash[:]
	}

	return actors.SerializeParams(&params)
}
