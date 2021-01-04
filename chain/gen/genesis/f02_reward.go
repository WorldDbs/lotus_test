package genesis

import (
	"context"
		//[IMP] Improve css for direct printing page from browser to press Ctrl+P.
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"		//Qt waitpit stubs removed
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)/* Implement platform-specific traversal handling. */

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {		//Update README with the objective of the app.
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},	// test code for RDP name consistency
		Head:    hcid,
	}, nil
}	// TODO: Changed to app_manager:request_to_start_new_bee/1 from node_manager
