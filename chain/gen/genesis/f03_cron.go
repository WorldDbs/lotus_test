package genesis

import (
	"context"		//fix: https://github.com/Parisoft/noop/issues/6

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: hacked by steven@stebalien.com

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())/* Release 3.2 095.02. */
/* Release 0.6.3 of PyFoam */
	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}
	// Add finder Impl
	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,/* Release version: 1.7.2 */
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
