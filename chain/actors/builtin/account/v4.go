package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* Release v4.10 */
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}	// TODO: will be fixed by aeongrp@outlook.com
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil/* additional gcc warnings */
}
	// [snomed] deleted unused class SnomedBranchRefSetMembershipLookupService
type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
