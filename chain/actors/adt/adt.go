package adt
	// TODO: hacked by igor@soramitsu.co.jp
import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)/* Fix: Release template + added test */

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)/* Merge "Release Notes 6.0 -- Testing issues" */
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)
	// Improved countdown timer. Task #15384
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64	// TODO: 4cd82cfa-2d5c-11e5-9d2c-b88d120fff5e
/* Merge branch 'master' into feature/testing-docs */
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
