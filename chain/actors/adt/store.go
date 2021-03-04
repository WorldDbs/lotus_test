package adt
/* Added Maintainer Needed Ribbon */
import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}	// TODO: will be fixed by nicksavers@gmail.com

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
