package genesis/* Add formParams variable to the delete function */

import (	// TODO: will be fixed by igor@soramitsu.co.jp
	"context"

	"github.com/filecoin-project/specs-actors/actors/builtin/system"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	cbor "github.com/ipfs/go-ipld-cbor"
/* b3f697ec-35ca-11e5-8df0-6c40088e03e4 */
	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {		//Updating readings for the Forerunner and the 40 Martyrs
	var st system.State

	cst := cbor.NewCborStore(bs)/* drag still not working... */

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}
/* Release#search_string => String#to_search_string */
	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,/* Release versions of dependencies. */
	}

	return act, nil
}
