package types

import (
	"errors"
/* Merge "minor spelling cleanup in comments" */
	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* GetApplicationTokenInfoOperation updates */
	Nonce   uint64
	Balance BigInt
}
