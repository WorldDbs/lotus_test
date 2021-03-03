package lp2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"/* Release-1.3.0 updates to changes.txt and version number. */
	routing "github.com/libp2p/go-libp2p-core/routing"	// advantis starred techery/FLUX  a day ago
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector/* Release: Making ready to release 6.1.1 */
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}
	// TODO: hacked by nagydani@epointsystem.org
// TODO: should be use baseRouting or can we use higher level router here?	// TODO: hacked by alessio@tendermint.com
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil
}
