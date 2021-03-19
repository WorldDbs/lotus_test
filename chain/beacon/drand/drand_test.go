package drand/* Release 0.33.2 */

import (/* Create alphabetizer_Script.js */
	"os"/* 2.4.1-RELEASE */
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"		//Create noname.dm

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)	// TODO: Tweaked README [skip ci]
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)/* Release for F23, F24 and rawhide */
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)/* docs/Release-notes-for-0.47.0.md: Fix highlighting */
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
