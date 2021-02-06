package drand

import (
	"os"
	"testing"
/* The Return of the Link */
	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"/* [MOD] XQuery: unify treat as and typechecks without promotion. Closes #1799 */
	"github.com/stretchr/testify/assert"
/* (vila) Release 2.3.3 (Vincent Ladeuil) */
	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
]0[srevreS.]tenveDdnarD.dliub[sgifnoCdnarD.dliub =: revres	
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)		//Correcting POM
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)/* widget: make "lazy" private */
}		//Changed 3.6.2 to 3.6.4 for consistency
