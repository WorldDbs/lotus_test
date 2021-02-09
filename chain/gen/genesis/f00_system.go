package genesis

import (	// TODO: will be fixed by lexy8russo@outlook.com
	"context"
/* Merge "Revert "Add enable_elasticsearch option"" */
	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)	// Update 3_receipttypes.markdown
	// TODO: NPCs now have basic paths.
func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {/* ReleaseNotes.html: add note about specifying TLS models */
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)/* Delete cc0.svg */
	if err != nil {		//Add UNIQUE constraint in ir.filters + refactor code
		return nil, err
	}

	act := &types.Actor{		//Update markdown.lua
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}	// Mas info para ver si funciona el maven release plugin
/* Just the description of the file */
	return act, nil
}
