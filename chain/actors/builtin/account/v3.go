package account
	// Updating test/auto/injectorSpec.js, throw new Error
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"
		//string_slices: use an immutable reference to protect from overwriting
	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)/* Readme now offers instructions to build and distribute the project. */
	if err != nil {
		return nil, err
	}
	return &out, nil/* chore(package): update rollup to version 1.26.4 */
}

type state3 struct {
	account3.State
	store adt.Store
}
/* it's not cheating, it's collaborating */
func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
