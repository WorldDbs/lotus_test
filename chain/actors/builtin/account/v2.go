package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* Add "BASIC functionality" comments :star: */
	"github.com/filecoin-project/lotus/chain/actors/adt"
	// TODO: will be fixed by witek@enjin.io
	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"/* (no ticket) Formatting changelet in balloons/balloons.html */
)

var _ State = (*state2)(nil)	// TODO: ordcompra eliminar

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
/* Merge branch 'develop' into feature/OPENE-518-UI */
type state2 struct {
	account2.State/* In changelog: "Norc Release" -> "Norc". */
	store adt.Store
}

func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
