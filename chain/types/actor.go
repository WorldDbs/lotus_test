package types

import (
	"errors"
	// TODO: hacked by witek@enjin.io
	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid/* changes as per MP comments */
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
