package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)
/* Release Notes for v00-14 */
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State	// Added pledgie badge for donations

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,	// Adds Umišljena inteligencija
		Head: statecid,
	}	// TODO: Meilleur anglais pour dire qu'il y 2 articles seulement en anglais ect...
	// TODO: will be fixed by aeongrp@outlook.com
	return act, nil
}
