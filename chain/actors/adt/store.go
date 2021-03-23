package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* Release 7.10.41 */
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore/* Release of eeacms/bise-backend:v10.0.26 */
}/* [MOD] add base controller */

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {	// TODO: calculated fields
	return adt.WrapStore(ctx, store)
}		//Update GDXProfiler.podspec
