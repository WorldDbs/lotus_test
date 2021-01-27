package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"		//add a solor test
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}/* debootstrap: upgrade to version 1.0.38 */

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
