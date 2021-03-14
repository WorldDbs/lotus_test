package lp2p

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)		//1de45908-2e41-11e5-9284-b827eb9e62be
}/* Update Release 8.1 */

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))/* StyleCop: Updated to support latest 4.4.0.12 Release Candidate. */
	return
}
