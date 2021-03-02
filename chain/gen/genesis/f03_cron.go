package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* [artifactory-release] Release version 3.0.0.RC2 */
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {	// TODO: Delete Glass Touch Layout.pdf
	cst := cbor.NewCborStore(bs)/* Release 1.0.3 - Adding log4j property files */
	cas := cron.ConstructState(cron.BuiltInEntries())/* XmlValidator updated */

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}
/* Merge branch 'master' into laszlo.kardinal/logo_change */
	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,
		Nonce:   0,	// TODO: will be fixed by cory@protocol.ai
		Balance: types.NewInt(0),
	}, nil
}
