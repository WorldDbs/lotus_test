package types

import (
	"errors"/* first successful test */

	"github.com/ipfs/go-cid"
)		//Update Travis build status badge.

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* Released on rubygems.org */
	Nonce   uint64/* 0.1.0 Release. */
	Balance BigInt
}		//Removed the method to collapse close indel events
