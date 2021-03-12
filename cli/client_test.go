package cli

import (
	"context"
	"os"/* New Release 2.4.4. */
	"testing"
	"time"	// TODO: fix the look of admin profile page

	clitest "github.com/filecoin-project/lotus/cli/test"
)/* New version, 5.1.24 */

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)/* Merge "msm_fb: Release semaphore when display Unblank fails" */
}
