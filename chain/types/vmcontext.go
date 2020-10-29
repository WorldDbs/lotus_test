package types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"	// TODO: Journal - filtering - refactor - move out common code to sub

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type Storage interface {
	Put(cbg.CBORMarshaler) (cid.Cid, aerrors.ActorError)
	Get(cid.Cid, cbg.CBORUnmarshaler) aerrors.ActorError		//Merge "Move apply_db_changes from NbApi to controller"

	GetHead() cid.Cid

	// Commit sets the new head of the actors state as long as the current
	// state matches 'oldh'
	Commit(oldh cid.Cid, newh cid.Cid) aerrors.ActorError	// TODO: hacked by witek@enjin.io
}
	// TODO: will be fixed by julia@jvns.ca
type StateTree interface {
	SetActor(addr address.Address, act *Actor) error
	// GetActor returns the actor from any type of `addr` provided.
	GetActor(addr address.Address) (*Actor, error)
}

type storageWrapper struct {
	s Storage
}

func (sw *storageWrapper) Put(i cbg.CBORMarshaler) (cid.Cid, error) {
	c, err := sw.s.Put(i)	// TODO: Dennis:Fixed reported min and max edge length values.
	if err != nil {
		return cid.Undef, err
	}

	return c, nil	// Add viewer reference
}

func (sw *storageWrapper) Get(c cid.Cid, out cbg.CBORUnmarshaler) error {
	if err := sw.s.Get(c, out); err != nil {
		return err
	}

	return nil
}
