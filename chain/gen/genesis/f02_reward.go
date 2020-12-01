package genesis

import (
	"context"		//Create SCCMCore.psm1
	// eb4e36a2-2e41-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}
/* add check for mw before checking config ( for pages wihout lib include ) */
	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* - Release to get a DOI */
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
