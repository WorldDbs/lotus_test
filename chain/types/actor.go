package types

import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* use always newest node v4.x version */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}/* Merge branch 'master' into PTX-1680 */
