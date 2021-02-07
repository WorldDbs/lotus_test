package types/* Release 0.45 */

import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {/* Release 0.6. */
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* TODOs before Release ergänzt */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
