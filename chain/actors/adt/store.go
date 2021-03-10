package adt
		//Update slitherhome.html
import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {/* [artifactory-release] Release version 0.7.5.RELEASE */
	Context() context.Context	// TODO: docs(README): typo CRA
	cbor.IpldStore
}
	// Context names fix
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* Merge branch 'master' into separate-note-switch */
	return adt.WrapStore(ctx, store)
}/* as pop3 bugs are fixed, it's time to remove workarounds */
