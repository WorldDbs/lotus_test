package genesis
/* 653a1706-2e60-11e5-9284-b827eb9e62be */
import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/util/adt"/* Finally released (Release: 0.8) */
	cbor "github.com/ipfs/go-ipld-cbor"/* Update ReleaseCycleProposal.md */

	bstore "github.com/filecoin-project/lotus/blockstore"	// TODO: hacked by praveen@minio.io
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}
	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	sms := market.ConstructState(a, h, h)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err/* Platform Release Notes for 6/7/16 */
	}

	act := &types.Actor{
		Code:    builtin.StorageMarketActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
