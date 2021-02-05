package cli

import (
	"context"
	"os"
	"testing"
	"time"
	// Added makepanda for building lui 
	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI/* AvatarDetails update */
// commands	// TODO: 79fd5716-2e6f-11e5-9284-b827eb9e62be
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()		//Macro module + module security
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)	// Add Português (Portugal)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
