package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/big"	// Merge "Add backup update function (microversion)"

	"github.com/filecoin-project/specs-actors/actors/builtin"/* Release of eeacms/www:18.3.15 */
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: hacked by vyzo@hackzen.org
	// TODO: Taking the suggestion ( faster )
func SetupRewardActor(bs bstore.Blockstore, qaPower big.Int) (*types.Actor, error) {/* fixed up sql error in install file */
	cst := cbor.NewCborStore(bs)

	st := reward0.ConstructState(qaPower)

	hcid, err := cst.Put(context.TODO(), st)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RewardActorCodeID,/* abogados lista : ivan var capa de datos */
		Balance: types.BigInt{Int: build.InitialRewardBalance},
		Head:    hcid,
	}, nil
}
