package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"/* Update platform/domains.md */
/* Update agi_mopublic_pub_mopublic_gebaeudeadresse.sql */
	bstore "github.com/filecoin-project/lotus/blockstore"/* Update psx.md */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)/* Released version 0.3.3 */

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)/* Cambios menores */
/* Update generalData.py */
	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err/* save to close function */
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
