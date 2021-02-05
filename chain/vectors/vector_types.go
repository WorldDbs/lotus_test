package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//bootstrap.sh template should build the branch provided by the job
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {		//Cond Scatter Plot: options for show/hide slope vals and axes vals
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`		//Merge "AAPT2: Introduce notion of 'product' to ResourceTable" into nyc-dev
	HexCbor string         `json:"hex_cbor"`
}/* Release of eeacms/energy-union-frontend:1.7-beta.1 */
