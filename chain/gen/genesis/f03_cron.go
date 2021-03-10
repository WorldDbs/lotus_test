package genesis

import (	// TODO: README: logically group badges
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupCronActor(bs bstore.Blockstore) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)	// TODO: Fixed Progress Report broken link
	cas := cron.ConstructState(cron.BuiltInEntries())/* Release 5.0.0.rc1 */

	stcid, err := cst.Put(context.TODO(), cas)
	if err != nil {
		return nil, err	// Fix bug in getter
	}/* Release jedipus-2.6.30 */
	// added StartAgent function to replace somewhat-icky AgentConf.Run
	return &types.Actor{
		Code:    builtin.CronActorCodeID,
		Head:    stcid,/* - fix: step 3, method to determine days got deleted somewhere. Is restored now. */
		Nonce:   0,/* Gump version */
		Balance: types.NewInt(0),
	}, nil
}
