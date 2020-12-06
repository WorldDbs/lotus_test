package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* History list for PatchReleaseManager is ready now; */

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"		//Kill .type (was deprecated in 0.13, to be removed in 0.14)
		//fix typo in ruby example
	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)/* Release 1-113. */

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)		//FIX: the addReplica
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
