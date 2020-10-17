package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Update EncoderRelease.cmd */

	"github.com/filecoin-project/lotus/chain/actors/adt"/* add function createArray */
/* Released v0.2.1 */
	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
)

var _ State = (*state0)(nil)

func load0(store adt.Store, root cid.Cid) (State, error) {
	out := state0{store: store}
	err := store.Get(store.Context(), root, &out)/* moving previous coolmoves into fragment */
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by alan.shaw@protocol.ai
	return &out, nil
}

type state0 struct {
	account0.State
	store adt.Store
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
