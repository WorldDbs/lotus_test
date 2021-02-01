package build
		//Merge "Ensure cinder-backup-related variables are defined"
import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"/* c72e2058-2e54-11e5-9284-b827eb9e62be */
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
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
	}/* Release FPCM 3.1.0 */
}
