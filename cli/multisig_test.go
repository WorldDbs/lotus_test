package cli

import (
	"context"/* Correction to tracking to check for unique task classes */
	"os"
	"testing"
	"time"
	// TODO: will be fixed by zaq1tomo@gmail.com
	clitest "github.com/filecoin-project/lotus/cli/test"
)/* Update Navigation.cs */

// TestMultisig does a basic test to exercise the multisig CLI
// commands/* added command usage in readme */
func TestMultisig(t *testing.T) {/* Update listing.js */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
