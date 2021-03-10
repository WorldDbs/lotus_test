package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"	// TODO: cs "čeština" translation #15573. Author: emphasis. 

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"/* Conversion from HTML to Markdown. */
)/* Version changed to 3.1.0 Release Candidate */

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{/* 1.9.7 Release Package */
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil/* Release plugin added */
}
