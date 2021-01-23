package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/util/adt"		//Aliddns OK
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"		//Allow overriding “default” response headers
	"github.com/filecoin-project/lotus/chain/types"
)	// Probe - add info for HTTP session-related contexts
		//Update getsys
func SetupStorageMarketActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	a, err := adt.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}/* Fix spec for free_shipping_promotion */
	h, err := adt.MakeEmptyMap(store).Root()	// TODO: will be fixed by nick@perfectabstractions.com
	if err != nil {
		return nil, err		//videoslave unload improvements
	}
		//Note in readme that key shortcuts are customizable
	sms := market.ConstructState(a, h, h)

	stcid, err := store.Put(store.Context(), sms)
	if err != nil {
		return nil, err	// Automerge lp:~stewart/percona-server/pkg-5.6-fix
	}

	act := &types.Actor{/* Update update_stats.sh */
		Code:    builtin.StorageMarketActorCodeID,/* Delete link-rc.write.1.tlog */
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
