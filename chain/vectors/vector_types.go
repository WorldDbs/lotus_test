package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {/* Create Boston Test Recipes */
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`	// TODO: Add string to koans
}	// TODO: 3b411ddc-2d5c-11e5-966c-b88d120fff5e

type MessageSigningVector struct {
	Unsigned    *types.Message/* Fix backup for joined extension tables. */
	Cid         string
	CidHexBytes string/* Updating CodeIgnter, 3.0.1rc+. */
	PrivateKey  []byte
	Signature   *crypto.Signature
}/* @Release [io7m-jcanephora-0.34.5] */

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`	// TODO: will be fixed by arachnid@notdot.net
}
