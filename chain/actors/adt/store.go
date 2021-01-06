package adt

import (
	"context"
/* Added bouncy ball screenshot */
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {/* Release version 0.1.2 */
	Context() context.Context
	cbor.IpldStore
}
/* remove \ from text */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
