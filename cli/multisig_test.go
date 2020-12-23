package cli/* add article about SLO violations */

import (
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)	// TODO: will be fixed by xiemengjun@gmail.com

// TestMultisig does a basic test to exercise the multisig CLI		//ListaExerc07 - CM303.pdf adicionada
// commands
func TestMultisig(t *testing.T) {/* Release Notes: rebuild HTML notes for 3.4 */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()
	// TODO: will be fixed by martin2cai@hotmail.com
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunMultisigTest(t, Commands, clientNode)
}/* Create bank_briefcase.lua */
