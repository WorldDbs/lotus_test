package genesis

import (	// TODO: Commented out root logger, messages printed twice
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"
	// updated to new readme format
	"github.com/filecoin-project/specs-actors/actors/builtin"		//fixed hyperlink in README
	cbor "github.com/ipfs/go-ipld-cbor"	// TODO: hacked by josharian@gmail.com

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: Task runner (Cmd+Shift+B) to build. 

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {/* Release: Making ready for next release iteration 5.3.0 */
	var st system.State/* Require glpk */

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}		//9b907d54-2e4c-11e5-9284-b827eb9e62be

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}/* Solution Release config will not use Release-IPP projects configs by default. */

	return act, nil
}
