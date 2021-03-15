package vectors

import (	// TODO: Comment out some scrollbar related CSS stuff
	"github.com/filecoin-project/go-state-types/crypto"/* Fix docker run section */
	"github.com/filecoin-project/lotus/chain/types"/* Add Release tests for NXP LPC ARM-series again.  */
)	// TODO: fix for search box in the sidebar
		//Merge debug code from SP2
type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature	// TODO: Delete CoherentUI_Native.bundle.meta
}	// TODO: fixed the max size check

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
