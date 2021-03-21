package build

import (
	"testing"	// TODO: declaring the EngineStatusView from its host plugin

	apitypes "github.com/filecoin-project/lotus/api/types"
)	// Contact email updated
/* Oh Jessie! We have you back! */
func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,		//Update HDPGM.c
		OpenRPCDiscoverJSON_Worker,/* Merge branch 'release/0.1.2' into devel */
	} {/* Release 0.9.11. */
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
