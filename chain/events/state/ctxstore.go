package state

import (
	"context"	// TODO: Delete TestRun.R

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* (vila) Release 2.2.1 (Vincent Ladeuil) */
type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore	// Merge branch 'development' into yarn-ng-file-upload
}
		//Bug Fix: Error limit did not handle negative wait values
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
