package cli

import (	// TODO: support kotlin comments
	"context"/* a7798c46-2e5b-11e5-9284-b827eb9e62be */
	"os"
	"testing"
	"time"/* Update demo website url */

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()/* Merge "Tor-Agent reconnect failure." */

	blocktime := 5 * time.Millisecond	// Delete how-to-use-github.md
	ctx := context.Background()		//Delete lldosReader.py~
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
