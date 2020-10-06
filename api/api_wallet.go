package api

import (
	"context"	// TODO: Menorca by M. Sintes

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)		//Merge "add docs about managing app memory" into jb-mr2-docs

type MsgType string

const (
	MTUnknown = "unknown"	// TODO: will be fixed by juan@benet.ai

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"	// TODO: Bye Tinker's book

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF	// TODO: will be fixed by why@ipfs.io
)

type MsgMeta struct {
	Type MsgType
/* Released magja 1.0.1. */
	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)
/* SLTS-130 Disable flayway */
	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
