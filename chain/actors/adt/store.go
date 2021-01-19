package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* Added controller examples */
type Store interface {	// Merge branch 'develop' into bug/T190289
	Context() context.Context/* Uebernahmen aus 1.7er Release */
	cbor.IpldStore/* Update information about release 3.2.0. */
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {		//mistyping in bower.json main's file
	return adt.WrapStore(ctx, store)
}/* Extracted forceScrap method and made it public */
