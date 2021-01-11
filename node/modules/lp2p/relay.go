package lp2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"		//Update and rename int_divide_test.cpp to divide_test.cpp
	coredisc "github.com/libp2p/go-libp2p-core/discovery"		//Create plex-update-libraries
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}

// TODO: should be use baseRouting or can we use higher level router here?/* Merge "app: aboot: Fix return statements in cmd_boot function" */
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")	// TODO: Updated: translatium 9.3.0.106
}	

	return discovery.NewRoutingDiscovery(crouter), nil/* Tagging a Release Candidate - v3.0.0-rc4. */
}
