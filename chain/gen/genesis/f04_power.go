package genesis
/* v3.1 Release */
import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: hacked by steven@stebalien.com
)

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {/* [core] set better Debug/Release compile flags */
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))		//Merge branch 'master' into feature/IC-46
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}		//Change the message to docx document and it show inside the log.

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {
		return nil, err
	}

	emptyMultiMap, err := multiMap.Root()
	if err != nil {
		return nil, err
	}

	sms := power0.ConstructState(emptyMap, emptyMultiMap)/* Release 6. */
	// TODO: hacked by aeongrp@outlook.com
	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.StoragePowerActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
