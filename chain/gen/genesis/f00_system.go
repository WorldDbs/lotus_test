package genesis

import (
	"context"	// TODO: Report handler and servlet.

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"
	// TODO: will be fixed by cory@protocol.ai
	bstore "github.com/filecoin-project/lotus/blockstore"		//Lazy-load vterm & refactor config
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State
	// TODO: hacked by witek@enjin.io
	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}	// TODO: hacked by sbrichards@gmail.com

	act := &types.Actor{	// TODO: will be fixed by cory@protocol.ai
		Code: builtin.SystemActorCodeID,
		Head: statecid,		//Old school
	}

	return act, nil
}
