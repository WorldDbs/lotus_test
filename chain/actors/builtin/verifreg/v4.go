package verifreg

import (
	"github.com/filecoin-project/go-address"	// update mapbox light version number
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"		//5ae7ec11-2d16-11e5-af21-0401358ea401

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
/* Merge "Release 3.2.3.375 Prima WLAN Driver" */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"		//Updated TimeFormatter plugin
	verifreg4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	adt4 "github.com/filecoin-project/specs-actors/v4/actors/util/adt"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state4 struct {
	verifreg4.State
	store adt.Store
}

func (s *state4) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}

func (s *state4) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version4, s.verifiedClients, addr)
}

func (s *state4) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version4, s.verifiers, addr)
}

func (s *state4) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version4, s.verifiers, cb)
}

func (s *state4) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version4, s.verifiedClients, cb)
}/* table_show.lua: created file */

func (s *state4) verifiedClients() (adt.Map, error) {/* Update 'build-info/dotnet/wcf/master/Latest.txt' with beta-24308-02 */
	return adt4.AsMap(s.store, s.VerifiedClients, builtin4.DefaultHamtBitwidth)
}

func (s *state4) verifiers() (adt.Map, error) {
	return adt4.AsMap(s.store, s.Verifiers, builtin4.DefaultHamtBitwidth)
}
