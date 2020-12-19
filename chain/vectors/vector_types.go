package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`	// TODO: Merge branch 'master' into frame-stats-in-global
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {/* Prepare of FreeDV 1.0.1 tag */
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
