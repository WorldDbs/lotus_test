package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* Release 2.1.0 */
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Modify celf to use initial cascade  */
	bstore "github.com/filecoin-project/lotus/blockstore"	// TODO: track if the app shutdown cleanly.
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {	// TODO: will be fixed by nicksavers@gmail.com
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* fix install issues */
		Balance: types.BigInt{Int: build.InitialRewardBalance},		//3.0.1 GM 4
		Head:    hcid,
	}, nil/* Release version: 2.0.5 [ci skip] */
}	// Fix of the last edited tab not being shown
