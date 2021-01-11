package account	// Merge branch 'devel' into dependabot/npm_and_yarn/mocha-8.4.0
/* Release: 5.7.3 changelog */
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* Release Candidate 0.5.6 RC6 */
	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {/* Remoevd unnecessary */
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil		//Merge branch 'service-vm-recovery' into authkeys_update
}

type state4 struct {
	account4.State
	store adt.Store
}
/* Fixes bold */
func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
