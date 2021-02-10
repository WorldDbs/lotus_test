package account

import (
	"github.com/filecoin-project/go-address"/* fix flake8 */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"		//Remove argument in output

	account2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
)

var _ State = (*state2)(nil)
/* Release 0.0.17 */
func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}	// TODO: Equals to NOTHINGNESS
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err/* Release note for #690 */
	}
	return &out, nil
}

type state2 struct {
	account2.State
	store adt.Store
}
/* thongtincanhan */
func (s *state2) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
