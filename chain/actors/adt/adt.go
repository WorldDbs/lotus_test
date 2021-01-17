package adt
	// TODO: added mappings function, added gsim and ddi 3.2
import (
	"github.com/ipfs/go-cid"
/* much better presentation of lang annotations in doc hover */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)
	// TODO: Update paladin.js
type Map interface {/* Release of eeacms/www:20.6.5 */
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
		//Improved speed of fp2_const_calc.
type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error		//add sample for assembler
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
