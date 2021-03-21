package cli

import (
	"context"/* * Alpha 3.3 Released */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)		//Update privilege.md

// TestClient does a basic test to exercise the client CLI
// commands
{ )T.gnitset* t(tneilCtseT cnuf
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
