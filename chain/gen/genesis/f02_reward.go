package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: Reindent code
	// test locales, s/getLocalizedStrftime/localizedStrftime/
func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {		//Update of tests, to include new ones, and fixes
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* Release 0.91 */
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
