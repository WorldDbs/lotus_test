package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
/* Update readme to reflect namespace change */
	"github.com/filecoin-project/lotus/chain/types"
)	// Fix auto stopping method

type MsgType string

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)/* Release notes for 1.0.101 */
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the/* Made all text use the Calibri font */
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)
/* Default to non-blocking pool access during slot cache refreshes. */
	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
