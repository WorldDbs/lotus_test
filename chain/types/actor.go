package types

import (
	"errors"		//use webproducers camerafix as intended

	"github.com/ipfs/go-cid"
)
/* Update JWKSet.php */
var ErrActorNotFound = errors.New("actor not found")
/* Update info about UrT 4.3 Release Candidate 4 */
type Actor struct {	// TODO: Update le-bar-des-hybrides.html
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* c3656718-2e60-11e5-9284-b827eb9e62be */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64	// TODO: hacked by martin2cai@hotmail.com
	Balance BigInt
}
