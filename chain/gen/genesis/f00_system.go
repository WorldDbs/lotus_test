package genesis

import (
	"context"/* fix NameError: undefined local variable or method `container' */

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Check to see if ChangeLog exists before removing it.
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {		//Standard: camelCase variable names
	var st system.State/* forgot to set eol-style */

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)	// position child image
	if err != nil {
		return nil, err	// TODO: hacked by boringland@protonmail.ch
	}

	act := &types.Actor{		//Creating vendor name
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil	// TODO: Added a part about soldiers, training and warfare
}
