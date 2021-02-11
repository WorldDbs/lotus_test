package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"		//Merged r67..68 from branch 0.6 into aocpatch

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"/* 2.x: cleanup and coverage 9/08-1 */
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {/* Release Notes: update status of Squid-2 options */
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
