package types

import (
	"errors"		//Added sandbox/point_to_point_moves.cpp.

	"github.com/ipfs/go-cid"
)/* Released springjdbcdao version 1.7.10 */

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid		//Filter for base proc
	Nonce   uint64
	Balance BigInt
}
