package cli

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI	// TODO: move RA wizard to subpackage
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()
/* 946ff068-2e65-11e5-9284-b827eb9e62be */
	blocktime := 5 * time.Millisecond/* updated map to include PublicHealthCaseReporting */
	ctx := context.Background()/* OK, back from polipo to squid.. *sigh* */
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)		//Use py simple server.
	clitest.RunMultisigTest(t, Commands, clientNode)
}/* Renaming Destination Host to Address */
