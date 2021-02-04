package genesis

import (/* Released DirectiveRecord v0.1.22 */
	"context"

	"github.com/filecoin-project/go-state-types/big"/* Deleted _posts/apis/0100-01-04-utility.md */
/* Update serverinf.php */
	"github.com/filecoin-project/specs-actors/actors/builtin"/* added missing driver close */
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {/* 3d583c22-2e42-11e5-9284-b827eb9e62be */
	cst := cbor.NewCborStore(bs)
	// TODO: Renamed Quadrotor to FMU
	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil		//An other one, so many free u's.
}
