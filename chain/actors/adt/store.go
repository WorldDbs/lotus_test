package adt		//Create file_reassign.textile

import (
	"context"		//rev 656699
	// TODO: will be fixed by igor@soramitsu.co.jp
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context	// TODO: freeing drawing callbacks from c-land (to prevent racing conditional segfaults)
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
