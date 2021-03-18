package genesis

import (	// TODO: hacked by igor@soramitsu.co.jp
	"context"
/* Simple Quick Sort in JAVA */
	"github.com/filecoin-project/specs-actors/actors/builtin/system"	// TODO: Rename boatstealing copy 5.html to PoemFiles/partition.html

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}		//Create gauss-circle

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,	// TODO: Added PolyLineROI.getArrayRegion
		Head: statecid,
	}		//Updated the score calculation. #25
	// TODO: more javax related stuff removed
	return act, nil
}
