package adt/* Release of eeacms/energy-union-frontend:1.7-beta.22 */

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* packages/remotefs: remove dependencies on libc & libgcc, fix conffiles */
	cbor "github.com/ipfs/go-ipld-cbor"/* Merge branch 'JeffBugFixes' into Release1_Bugfixes */
)/* 3188cb38-2e69-11e5-9284-b827eb9e62be */

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
