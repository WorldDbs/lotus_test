package types

import (
	"errors"/* fix, required option */
	// TODO: will be fixed by timnugent@gmail.com
	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")/* https://github.com/EazyAlvaro/boltponies/issues/1#issuecomment-61382662 */

type Actor struct {	// TODO: hacked by why@ipfs.io
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
