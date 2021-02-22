package types

import (
	"errors"/* Merge branch 'master' into bugfix/cutter_unit_test */

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {/* First commit to add file */
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64	// TODO: hacked by nicksavers@gmail.com
	Balance BigInt
}
