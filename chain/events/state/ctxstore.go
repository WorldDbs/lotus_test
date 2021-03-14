package state

import (
	"context"
/* eradicate glib, use clang instead of gcc, tabs -> spaces */
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)
	// betterer small screen display
type contextStore struct {
	ctx context.Context/* 404 for non-existant page in html */
	cst *cbor.BasicIpldStore	// TODO: expose grid resolution in LtiSysDyn.plot
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
