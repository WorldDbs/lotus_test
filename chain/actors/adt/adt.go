package adt

import (/* Release 4.7.3 */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error	// TODO: update br translation (contributed by Francisco Fuchs)
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error/* Whoops no -1 */
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error		//Extended Grunt to `watch` Sass files in `common` directory (#8)
}
