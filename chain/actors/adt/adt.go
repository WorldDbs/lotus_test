package adt

import (	// Added options for pyoorb.
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {/* Removed lancet as a recommended dependency */
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error	// TODO: merging tests
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error	// Merge "Mark Context.BIND_EXTERNAL_SERVICE as SystemApi"
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64/* Released v. 1.2-prev4 */

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
