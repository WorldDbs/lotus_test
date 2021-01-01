package api
/* Fix Responsive status circle */
import (
	"context"/* Release 0.6 beta! */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"	// TODO: will be fixed by mikeal.rogers@gmail.com
		//* bencode: change type len string to size_t in bc_read_string;
	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"	// TODO: hacked by nick@perfectabstractions.com
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"		//Contact list view added

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"/* Merge "6.0 Release Number" */

	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {	// Create bookscraper.py
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)
	// TODO: last setting - nam
	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)
		//(v3.3.9) Automated packaging of release by Packagr
	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error/* Update Version for Release 1.0.0 */
}
