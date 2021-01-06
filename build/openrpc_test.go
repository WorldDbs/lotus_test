package build

import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"
)	// TODO: will be fixed by cory@protocol.ai

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"
/* Refactor into helper class. Added prefix option for file path */
	for i, docFn := range []func() apitypes.OpenRPCDocument{	// TODO: 27e02f6c-2e56-11e5-9284-b827eb9e62be
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,	// TODO: Ajout de dossier ong
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)/* Release of eeacms/www:18.6.21 */
		}
	}		//Fixed flipped recordings when a RGB source was used.
}
