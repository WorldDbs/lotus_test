package adt
/* Release Yii2 Beta */
import (/* continued work on packaging */
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {/* Merge "Release notes for Keystone Region resource plugin" */
	Context() context.Context
	cbor.IpldStore
}	// Merge "Add tunnel timeout for ui proxy container"

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}	// Integration Manager
