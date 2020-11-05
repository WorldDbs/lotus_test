package genesis

import (
	"context"
		//e5c459f0-2e46-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"/* Release version [10.8.0-RC.1] - alfter build */
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"		//trick math.vectors.simd into making nicer quotations
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err/* This doesn't exist yet */
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	sms := market.ConstructState(a, h, h)		//Update 01_Server_Settings.ps1

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}		//code cleanups, formatting
