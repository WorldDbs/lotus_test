package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)
	// TODO: will be fixed by mowrain@yandex.com
func load3(store adt.Store, root cid.Cid) (State, error) {/* Rename Bhaskara.exe.config to bin/Release/Bhaskara.exe.config */
	out := state3{store: store}	// TODO: anticipate initial data and tests
	err := store.Get(store.Context(), root, &out)
	if err != nil {		//Add probably useless _printf(x,y) macro
		return nil, err
	}
	return &out, nil
}
/* Fix phpunit compatibility */
type state3 struct {
	account3.State
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
