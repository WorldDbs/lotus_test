package account/* Release of stats_package_syntax_file_generator gem */

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

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
	return &out, nil/* Merge "Adding job_execution_update api call" */
}

type state0 struct {
	account0.State
	store adt.Store	// Merge "Correct rabbit messaging config set in devstack"
}

func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
