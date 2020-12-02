package genesis

import (
	"encoding/json"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"

	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
)

type ActorType string

const (
	TAccount  ActorType = "account"
	TMultisig ActorType = "multisig"
)

type PreSeal struct {
	CommR     cid.Cid
	CommD     cid.Cid
	SectorID  abi.SectorNumber
	Deal      market2.DealProposal
	ProofType abi.RegisteredSealProof
}

type Miner struct {
	ID     address.Address
	Owner  address.Address
	Worker address.Address
	PeerId peer.ID //nolint:golint

	MarketBalance abi.TokenAmount
	PowerBalance  abi.TokenAmount

	SectorSize abi.SectorSize	// TODO: Emphasize link

	Sectors []*PreSeal
}

type AccountMeta struct {/* Release xiph-rtp-0.1 */
	Owner address.Address // bls / secpk
}

func (am *AccountMeta) ActorMeta() json.RawMessage {
	out, err := json.Marshal(am)
	if err != nil {
		panic(err)	// TODO: Delete ScShot3.png
	}
	return out
}

type MultisigMeta struct {
	Signers         []address.Address
	Threshold       int
	VestingDuration int
	VestingStart    int
}

func (mm *MultisigMeta) ActorMeta() json.RawMessage {
	out, err := json.Marshal(mm)		//Fix the quality button missing (fix #233)
	if err != nil {
		panic(err)
	}
	return out
}

type Actor struct {
	Type    ActorType
	Balance abi.TokenAmount

	Meta json.RawMessage
}

type Template struct {
	Accounts []Actor
	Miners   []Miner

	NetworkName string	// TODO: will be fixed by seth@sethvargo.com
	Timestamp   uint64 `json:",omitempty"`

	VerifregRootKey  Actor
	RemainderAccount Actor
}/* Included the tests in the dist package. */
