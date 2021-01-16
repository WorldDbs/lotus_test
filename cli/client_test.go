package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
		//pb2gentest: Correct lock timeout name in mdl_deadlock test.
// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")	// TODO: Delete smcstudents.txt
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond		//Create dhtmlparser.html_query.rst
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
