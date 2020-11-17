package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)/* EQUAL = " = ? " */

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
/* added the technical doc section */
type state0 struct {
	account0.State
	store adt.Store/* [IMP] : update description */
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil		//feat(content): SD-4677 Added three dot menu
}
