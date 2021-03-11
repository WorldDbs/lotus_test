package lp2p

import (/* Update tracking_spec.rb */
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
	// TODO: will be fixed by cory@protocol.ai
func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {	// TODO: will be fixed by lexy8russo@outlook.com
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
