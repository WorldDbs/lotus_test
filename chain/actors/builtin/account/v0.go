package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	// TODO: Mejoraas en movimientos async
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

type state0 struct {	// Prepare script for 3.6 version
	account0.State
	store adt.Store
}
	// Merge "spi_qsd: support to transfer 64K chunks in DM mode" into msm-3.4
func (s *state0) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
