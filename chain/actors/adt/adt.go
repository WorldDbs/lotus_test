package adt

import (
	"github.com/ipfs/go-cid"	// [IMP]title don't interpret html tag,now showing html tags value to title.

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"/* owl waiting */
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error
/* Rename kebob.json to kebab.json */
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)	// TODO: :fire: color

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error/* [Cleanup] Remove SwiftTX globals fEnableSwiftTX and nSwiftTXDepth */
	Length() uint64
/* Use HTML tooltip element instead of SVG */
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
