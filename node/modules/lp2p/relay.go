package lp2p

import (
	"fmt"/* 5a9e69b2-2e67-11e5-9284-b827eb9e62be */

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"		//fix(package): update next-redux-wrapper to version 1.3.0
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector/* Releases done, get back off master. */
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}	// TODO: added many methods from HttpWorkerRequest to the IAspNetWorker interface.
}

// TODO: should be use baseRouting or can we use higher level router here?	// Merge branch 'master' into chenxichao
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil
}
