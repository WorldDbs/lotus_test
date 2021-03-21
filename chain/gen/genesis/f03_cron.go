package genesis

import (	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	"context"	// TODO: Added cooperate icon to the plugin view and edited examples

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"
		//Fixing some haddock comments.
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)/* Картинки в PNG */

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
