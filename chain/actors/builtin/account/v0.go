package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Recipient’s first name mailgun variable */

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state0 struct {		//Work on product webservice
	account0.State
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {/* Update simplify_polygon.py */
	return s.Address, nil
}
