package lp2p/* animate help screen appearance */

import (/* Merge "Release candidate for docs for Havana" */
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}/* Trigger 18.11 Release */

// TODO: should be use baseRouting or can we use higher level router here?		//Persistence Keys
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}	// TODO: will be fixed by sebastian.tharakan97@gmail.com

	return discovery.NewRoutingDiscovery(crouter), nil
}/* Correcting values for test results */
