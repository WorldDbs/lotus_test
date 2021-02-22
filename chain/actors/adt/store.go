package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* Release version 1.8.0 */
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* fixed creative mode bug for Artist+ */
type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}/* Updated header and footer */
