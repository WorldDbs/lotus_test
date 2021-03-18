package account

import (
	"github.com/filecoin-project/go-address"	// Fix the disappearing image problem
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

)lin()3etats*( = etatS _ rav

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}		//Update alexandre.html
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
/* 0.1.2 Release */
type state3 struct {/* convert int to str */
	account3.State
	store adt.Store
}
	// TODO: Add build and coverage status
func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
