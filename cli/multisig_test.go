package cli	// iteration.

import (/* Preparations to add incrementSnapshotVersionAfterRelease functionality */
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond/* b750c02e-2e53-11e5-9284-b827eb9e62be */
	ctx := context.Background()/* Initial Release - Supports only Wind Symphony */
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}		//9d5a3f06-2e5d-11e5-9284-b827eb9e62be
