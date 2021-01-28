package cli	// Update/Create boMAoMmXlZGwGJcDbgCk9w_img_0.jpg

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"/* Merge "Release 1.1.0" */
)
	// Use a variable to explicitly trust global config files
// TestClient does a basic test to exercise the client CLI
// commands	// TODO: Updated test cases and psr2 fixes
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* Release Candidate 0.5.8 RC1 */
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
