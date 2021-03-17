package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"/* handle case where no logistics options exist for an orer form */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"/* more concrete class-loader dependency */
)
/* 54675ac2-2e62-11e5-9284-b827eb9e62be */
func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)/* Update to statinamic 0.4.2 */
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
