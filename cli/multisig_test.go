package cli		//do not show root partition in debug mode
	// TODO: StringUtil whitespace clean
import (/* Merge "Release notes for Danube 2.0" */
	"context"/* Add Release-Notes for PyFoam 0.6.3 as Markdown */
	"os"
	"testing"
	"time"	// TODO: Merge branch 'develop' into fix/test_robustness
	// TODO: will be fixed by nicksavers@gmail.com
	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI
// commands
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
