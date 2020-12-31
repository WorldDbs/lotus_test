package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"		//install create-new-uuid in setup.py

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)
/* Implemented all missing placeholder application server tests. */
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err	// TODO: Revert accidentally merged code
	}/* Added pyexiftool */

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}
