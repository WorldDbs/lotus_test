package cli

import (
	"context"
	"os"	// TODO: Delete ProductosVista.php
	"testing"
"emit"	

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)/* Update vaadin dependency to 7.7.2 */
	clitest.RunMultisigTest(t, Commands, clientNode)
}
