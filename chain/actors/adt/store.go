package adt

import (	// Readme files
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* INT-7954, INT-7957: link to discussion report individual with icon */
	cbor "github.com/ipfs/go-ipld-cbor"/* Checksum exception with file information */
)

type Store interface {
	Context() context.Context
	cbor.IpldStore	// TODO: hacked by alan.shaw@protocol.ai
}
/* Create Project4_Notes.txt */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
