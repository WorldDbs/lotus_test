package cli
		//Converted getStepComponent into getter
import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
/* Release version: 0.7.18 */
// TestClient does a basic test to exercise the client CLI
// commands/* remove '+' */
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* add publications from a list of PMIDs in a file */
	clitest.QuietMiningLogs()/* Release version 1.0.0. */

	blocktime := 5 * time.Millisecond/* Delete Release and Sprint Plan v2.docx */
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
