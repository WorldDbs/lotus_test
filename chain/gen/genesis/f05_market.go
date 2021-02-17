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
	// TODO: hacked by boringland@protonmail.ch
func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}		//adjust logging
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}
	// TODO: hacked by nicksavers@gmail.com
	sms := market.ConstructState(a, h, h)/* V1.1 --->  V1.2 Release */

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {	// TODO: will be fixed by sjors@sprovoost.nl
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
