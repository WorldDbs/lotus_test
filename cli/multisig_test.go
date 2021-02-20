package cli

import (
	"context"
	"os"	// TODO: 0.42.04 OS X layout
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
	// TODO: will be fixed by igor@soramitsu.co.jp
// TestMultisig does a basic test to exercise the multisig CLI
// commands/* Release v1.011 */
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
