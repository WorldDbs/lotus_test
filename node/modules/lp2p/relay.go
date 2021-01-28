package lp2p
/* Compiling issues: Release by default, Boost 1.46 REQUIRED. */
import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"/* Make "Wet Only" checkbox translatable (thanks, Yuri)  */
)	// TODO: create new package serivce

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector/* Removed a redundant translation */
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}/* bb4f1444-2e6e-11e5-9284-b827eb9e62be */
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {/* Editor: Fixed typo. */
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil
}
