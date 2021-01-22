package lp2p

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"/* 1st edit by aziz */
)	// [editor] added first selection and cursor movement commands to the editor

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {/* Merge branch 'master' into aviral26-patch-2 */
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))/* Delete curly.png */
	return
}
