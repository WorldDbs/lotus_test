package account/* Added 'the most important changes since 0.6.1' in Release_notes.txt */
/* Release 1.3.10 */
import (/* fix(deps): update dependency polished to v3.0.3 */
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
)

var _ State = (*state3)(nil)

func load3(store adt.Store, root cid.Cid) (State, error) {
	out := state3{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {/* Release 11.1 */
		return nil, err/* 6cf751d2-2e53-11e5-9284-b827eb9e62be */
	}
	return &out, nil
}

type state3 struct {
	account3.State
	store adt.Store
}

func (s *state3) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}
