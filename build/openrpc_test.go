package build

import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"		//Make style/index.js convention work for prebuilt extensions.
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {		//Create 02.NumbersEndingIn7.java
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
