package account
	// TODO: Laravel 5.2 availability
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	account4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
)

var _ State = (*state4)(nil)/* Versions managed in separated class */

func load4(store adt.Store, root cid.Cid) (State, error) {/* Release for another new ESAPI Contrib */
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {/* 2aab957e-2e65-11e5-9284-b827eb9e62be */
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	account4.State
	store adt.Store
}

func (s *state4) PubkeyAddress() (address.Address, error) {
	return s.Address, nil
}		//FB2 Output: Support SVG images in the input document
