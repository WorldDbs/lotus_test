package cli/* more example info */

import (
	"context"/* Released reLexer.js v0.1.1 */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {/* [artifactory-release] Release version 0.9.15.RELEASE */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond/* b24cff1c-2e66-11e5-9284-b827eb9e62be */
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
