package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"
		//Merge "COMP: Add images to the MultiplyTwoImages documentation"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Create chasing summer 1.html */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)		//69a75764-2e6e-11e5-9284-b827eb9e62be
	if err != nil {
		return nil, err
	}
	// TODO: will be fixed by timnugent@gmail.com
	return &types.Actor{/* [Merge]: merge with addons2 */
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,/* CWS-TOOLING: rebase CWS printerpullpages to trunk@270723 (milestone: DEV300:m46) */
	}, nil
}
