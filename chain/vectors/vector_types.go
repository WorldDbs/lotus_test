package vectors
	// TODO: will be fixed by lexy8russo@outlook.com
import (
	"github.com/filecoin-project/go-state-types/crypto"/* Released version 0.3.3 */
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {	// Arc Widget: Handle legacy color, fill
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
}/* controllers/values: emit{Change -> Updated} */
	// TODO: will be fixed by timnugent@gmail.com
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}/* Release of eeacms/www-devel:19.10.22 */
