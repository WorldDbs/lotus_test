package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"	// Attempt to handle other axes
)
/* Documentation technique */
// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {	// TODO: will be fixed by fjl@ethereum.org
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
