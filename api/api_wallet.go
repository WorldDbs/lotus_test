package api
		//Minimax speedup, call get_legal_moves later
import (/* Release 0.11.2. Review fixes. */
	"context"		//Change license in commands.h

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string/* Release 1.2.0.13 */

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes/* Merge "Migration to using requests" */
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)/* Initial Release of the README file */
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)/* getBoard(String[] a) */

type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the	// TODO: will be fixed by hugomrdias@gmail.com
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)
	// update npm config to spec, and semver range specifier to ^
	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}
