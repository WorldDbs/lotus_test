package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)/* Released 8.0 */
/* Log response message for invalid token errors. */
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
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
