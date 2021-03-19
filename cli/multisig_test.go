package cli
/* initial support for server settings added to rest facade */
import (
	"context"
	"os"
	"testing"/* Merge branch 'master' into Release-5.4.0 */
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestMultisig does a basic test to exercise the multisig CLI		//Bug Fixed and Set Rank Added :)
// commands	// Create AnvilDamageInfo.java
func TestMultisig(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()/* Release 0.95.040 */
	// TODO: Fixed some files, renamed Versions folder to dist
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}
