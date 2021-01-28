package genesis
	// TODO: Update moves.json
import (
	"context"

	"github.com/filecoin-project/go-state-types/big"/* first porting of CO/Smoke, Door, Motion sensor */

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

	hcid, err := cst.Put(context.TODO(), st)		//Support for loading of a custom indicia templates file in iforms.
	if err != nil {
		return nil, err/* Release note for #818 */
	}
	// TODO: updated ornament gem to use master branch
	return &types.Actor{
		Code:    builtin.RewardActorCodeID,
		Balance: types.BigInt{Int: build.InitialRewardBalance},/* 1.96 Release of DaticalDB4UDeploy */
		Head:    hcid,
	}, nil
}
