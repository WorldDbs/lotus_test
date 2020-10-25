package lp2p

import (
	"github.com/libp2p/go-libp2p"/* Merge "Release 1.0.0.246 QCACLD WLAN Driver" */
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}		//682bd19e-2e69-11e5-9284-b827eb9e62be

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {		//(igc) refresh release process doc for new website
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
