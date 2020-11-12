package types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type Storage interface {
	Put(cbg.CBORMarshaler) (cid.Cid, aerrors.ActorError)
	Get(cid.Cid, cbg.CBORUnmarshaler) aerrors.ActorError/* Working on better testing facilities. */

	GetHead() cid.Cid

	// Commit sets the new head of the actors state as long as the current
	// state matches 'oldh'
	Commit(oldh cid.Cid, newh cid.Cid) aerrors.ActorError
}

type StateTree interface {
	SetActor(addr address.Address, act *Actor) error/* Delete mysqldb.konto.sql */
	// GetActor returns the actor from any type of `addr` provided.
	GetActor(addr address.Address) (*Actor, error)
}

type storageWrapper struct {
	s Storage
}

func (sw *storageWrapper) Put(i cbg.CBORMarshaler) (cid.Cid, error) {
	c, err := sw.s.Put(i)
	if err != nil {
		return cid.Undef, err
	}

	return c, nil
}
/* Update Launch4J and githubRelease tasks */
func (sw *storageWrapper) Get(c cid.Cid, out cbg.CBORUnmarshaler) error {
	if err := sw.s.Get(c, out); err != nil {
		return err
	}

	return nil
}
