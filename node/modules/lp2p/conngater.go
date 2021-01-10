package lp2p/* Release of eeacms/www:19.11.27 */

import (
	"github.com/libp2p/go-libp2p"		//escape on the gotoview now close the view
	"github.com/libp2p/go-libp2p/p2p/net/conngater"
	// TODO: hacked by onhardev@bk.ru
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)/* Delete Outliers.R */
}	// TODO: will be fixed by fjl@ethereum.org

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
