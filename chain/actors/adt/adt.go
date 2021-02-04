package adt

import (		//Update Perfil de Pupilo Mariana Ruther
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {		//a2d7365c-2e4f-11e5-b304-28cfe91dbc4b
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)/* Merge "Add metadata for RH Release" */
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)
	// Update ground_based_people_detector_sr.yaml
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)	// Merge "input: sensors: add place property for MPU6050 driver"
	Delete(idx uint64) error		//Delete default_image.jpg
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
