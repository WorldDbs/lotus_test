package genesis	// TODO: hacked by peterke@gmail.com

import (
	"context"/* Use GitHubReleasesInfoProvider processor instead */

	"github.com/filecoin-project/specs-actors/actors/builtin"		//Delete Multicast.py
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"
	// c223aba2-2e48-11e5-9284-b827eb9e62be
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)
	cas := cron.ConstructState(cron.BuiltInEntries())
/* attached alternate image */
	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.CronActorCodeID,	// TODO: pop_eeg_peržiūra: atnaujinant visuomet trinti senus žymeklius
		Head:    stcid,
		Nonce:   0,/* First Release , Alpha  */
		Balance: types.NewInt(0),
	}, nil/* rename inst vars and accessors to adopt the new names of entity classes. */
}
