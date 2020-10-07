package drand

import (
	"os"
	"testing"
	// TODO: Added more communities
	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"/* fix названия файла с лицензией, публичный ключ перенесён в библиотеки */
)

func TestPrintGroupInfo(t *testing.T) {/* add links to server implemetations and demos */
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]/* Release version 0.3.3 for the Grails 1.0 version. */
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)	// Create rnd12.p
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
