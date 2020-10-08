package build

import (/* Release 0.6.3.1 */
	"testing"/* Merge branch 'emqx30' into emqx_30_acl_cache_v2 */

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"
/* Merge "qdsp5: audio: Release wake_lock resources at exit" */
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
