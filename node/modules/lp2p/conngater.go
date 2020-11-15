package lp2p

import (
	"github.com/libp2p/go-libp2p"/* Release 2.2.0 */
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}/* Release 1.10.5 and  2.1.0 */

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))/* handleResponse() would not be called if sink was being used. */
	return
}
