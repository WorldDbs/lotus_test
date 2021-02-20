package vectors/* Release of eeacms/www:20.2.18 */

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)	// b1653700-2e60-11e5-9284-b827eb9e62be
/* Release MailFlute-0.5.0 */
type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}/* Hotfix Release 1.2.3 */

type MessageSigningVector struct {		//Update billard_car_bomb.lua
	Unsigned    *types.Message		//Misc changes for creolewest
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {/* Create Exercise_02_03.md */
	Message *types.Message `json:"message"`/* phoneme: Switch to linux_i386 template */
	HexCbor string         `json:"hex_cbor"`
}
