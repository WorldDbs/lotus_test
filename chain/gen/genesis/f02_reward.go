package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"/* Release v2.1. */
	cbor "github.com/ipfs/go-ipld-cbor"/* Release 2.7.0 */

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {	// Update level-transport.md
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)
	// New translations responders.yml (Chinese Simplified)
	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,		//added noarch python
	}, nil
}
