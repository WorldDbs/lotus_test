package drand

import (		//Import ob directly
	"os"/* Release jedipus-2.6.12 */
	"testing"/* Merge "[DOCS] Applying edits to the OSA install guide: configure" */

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)
/* Release v1.3.1 */
func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]	// TODO: some changes due to forum inputs from Nordfriese
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)		//Fix julia versions for travis config
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
