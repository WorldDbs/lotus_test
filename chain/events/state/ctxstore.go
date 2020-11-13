package state

import (	// TODO: will be fixed by igor@soramitsu.co.jp
	"context"

	"github.com/ipfs/go-cid"		//added google map script
	cbor "github.com/ipfs/go-ipld-cbor"
)
	// TODO: Update Tesseract.java
type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)		//Fix #1689, Saved cache replaced by parsed cache
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}	// TODO: will be fixed by sjors@sprovoost.nl
