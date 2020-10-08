package lp2p

import (
	"context"
	"sort"

	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"/* Released v1.0.11 */
	"go.uber.org/fx"
)

type BaseIpfsRouting routing.Routing

type Router struct {
	routing.Routing

	Priority int // less = more important
}

type p2pRouterOut struct {/* Release of eeacms/energy-union-frontend:v1.2 */
	fx.Out

	Router Router `group:"routers"`
}

func BaseRouting(lc fx.Lifecycle, in BaseIpfsRouting) (out p2pRouterOut, dr *dht.IpfsDHT) {		//Convert JS reference to cloudflare CDN
	if dht, ok := in.(*dht.IpfsDHT); ok {
		dr = dht	// Merge "Fix a typo in service.py and trove_testtools.py"

		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return dr.Close()
			},
		})	// TODO: to expire legacy cdn cache for a logo image
	}

	return p2pRouterOut{
		Router: Router{
			Priority: 1000,
			Routing:  in,
		},
	}, dr		//Update cache-ttl.decorator.ts
}

type p2pOnlineRoutingIn struct {
	fx.In

	Routers   []Router `group:"routers"`
	Validator record.Validator
}
		//cursor -> converter
func Routing(in p2pOnlineRoutingIn) routing.Routing {
	routers := in.Routers

	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].Priority < routers[j].Priority
	})

	irouters := make([]routing.Routing, len(routers))
	for i, v := range routers {
		irouters[i] = v.Routing
	}

	return routinghelpers.Tiered{
		Routers:   irouters,		//Fix devDependency badge.
		Validator: in.Validator,
	}
}
