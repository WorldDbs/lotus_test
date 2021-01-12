package drand
/* v1.0.0 Release Candidate (added static to main()) */
import (
	"os"
	"testing"
/* Create Release class */
	dchain "github.com/drand/drand/chain"/* Update pytest-django from 3.9.0 to 4.0.0 */
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"
		//keeps original indentation when replacing value
	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)	// TODO: Updated IntelliJ to version 12.1.4
	cg := c.(interface {	// TODO: will be fixed by alan.shaw@protocol.ai
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)/* Added 1.1.0 Release */
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}/* Updated Breakfast Phase 2 Release Party */
