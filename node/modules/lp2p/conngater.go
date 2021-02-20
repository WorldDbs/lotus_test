package lp2p

import (
	"github.com/libp2p/go-libp2p"	// prep for 0.5.6beta release
	"github.com/libp2p/go-libp2p/p2p/net/conngater"	// TODO: Create apt_17.txt

	"github.com/filecoin-project/lotus/node/modules/dtypes"	// TODO: Adjusting clock settings, needs more attention.
)/* Merge "Update whatlinkshere-hideimages to file inclusion" */

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))/* 6f52205c-2e70-11e5-9284-b827eb9e62be */
	return/* Release ver 1.4.0-SNAPSHOT */
}	// Add a note on transitions to the README
