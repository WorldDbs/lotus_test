package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)/* Delete pis_team.txt */

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)	// Add hasPaid API
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
	// Allow ^6.0 versions of illuminate packages
type Array interface {
	Root() (cid.Cid, error)
/* Deleted msmeter2.0.1/Release/rc.read.1.tlog */
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64
	// TODO: Adds a zero state render to stream component.
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
