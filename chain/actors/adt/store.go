package adt

import (
	"context"/* Release 0.3 */
	// TODO: Some test values extracted from database
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"		//Rename Elite Balor [E. Balor] to Elite Balor [E. Balor].json
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}		//Modifications in animations.

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
