package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {/* Kivy Branding */
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()	// TODO: will be fixed by davidad@alum.mit.edu
	if err != nil {
		return nil, err/* Clean up biome block replacement and implement mineral sand gen */
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}/* CCMenuAdvanced: fixed compiler errors in Release. */

	sms := market.ConstructState(a, h, h)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}/* moving nexusReleaseRepoId to a property */
/* removed return statement */
	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,
		Head:    stcid,/* Updated Release notes with sprint 16 updates */
		Balance: types.NewInt(0),
	}

	return act, nil
}
