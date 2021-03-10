package build

import (
	"testing"/* Merge "resourceloader: Release saveFileDependencies() lock on rollback" */

	apitypes "github.com/filecoin-project/lotus/api/types"
)
	// TODO: will be fixed by nagydani@epointsystem.org
func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"
	// TODO: python/build/libs.py: update FFmpeg to 4.2.2
	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)		//Persists buildTasks in the indexedDB.
		}
	}
}
