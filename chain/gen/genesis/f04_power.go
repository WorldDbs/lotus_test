package genesis

import (
	"context"
	// Smoother SPEC
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/util/adt"/* Release 0.97 */

	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"/* Release of eeacms/www-devel:18.6.12 */
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStoragePowerActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {
		return nil, err
	}

	emptyMultiMap, err := multiMap.Root()
	if err != nil {
		return nil, err/* Rename fun-with-docker to fun-with-docker.txt */
	}

	sms := power0.ConstructState(emptyMap, emptyMultiMap)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.StoragePowerActorCodeID,/* Release areca-5.4 */
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
