package vectors/* Release 1.34 */

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"/* 961254f2-2e64-11e5-9284-b827eb9e62be */
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`	// Invert warning checking.
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
/* Add IX-F ID 807 for GABIX */
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
