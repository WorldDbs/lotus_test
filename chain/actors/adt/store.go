package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context	// cleanup some formatting of nbttools class
	cbor.IpldStore
}
/* preparing release 3.6 */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)/* Prepare to Release */
}	// TODO: will be fixed by alan.shaw@protocol.ai
