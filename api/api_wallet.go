package api

import (
	"context"/* integrated callback functions in start page */

	"github.com/filecoin-project/go-address"/* Merge "Release 5.4.0" */
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string		//Update game.info

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"	// TODO: 7862086a-2e52-11e5-9284-b827eb9e62be

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

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

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)		//Merge branch 'master' into strictFunctionTypes
	WalletDelete(context.Context, address.Address) error/* default cookie path is now "/" */
}
