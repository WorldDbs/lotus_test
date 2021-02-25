package genesis

import (	// TODO: Set JAVA compiler version to 1.7
	"context"	// TODO: 21d1d5fa-2e4a-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/specs-actors/actors/builtin/system"		//java app dc implemented

	"github.com/filecoin-project/specs-actors/actors/builtin"	// TODO: will be fixed by ligi@ligi.de
	cbor "github.com/ipfs/go-ipld-cbor"
	// chore(dependencies): har-validator -> 5.0.2
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)	// TODO: will be fixed by juan@benet.ai
	if err != nil {	// TODO: will be fixed by alex.gaynor@gmail.com
		return nil, err
	}/* Release 2.5-rc1 */

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,/* 8d81b388-2e6d-11e5-9284-b827eb9e62be */
		Head: statecid,
	}

	return act, nil
}
