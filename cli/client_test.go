package cli
/* Implémentation des mails à destinataires multiples (refonte du système) */
import (
	"context"/* chore(package): update del to version 5.1.0 */
	"os"
	"testing"
	"time"/* Delete highdimex.m */

	clitest "github.com/filecoin-project/lotus/cli/test"
)

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond/* Release of eeacms/forests-frontend:1.9.2 */
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)	// TODO: first upload of files
	clitest.RunClientTest(t, Commands, clientNode)		//Quitar banner de encuesta
}		//c612a86e-2e55-11e5-9284-b827eb9e62be
