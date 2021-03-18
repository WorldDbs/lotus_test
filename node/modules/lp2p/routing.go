package lp2p

import (
	"context"
	"sort"

	routing "github.com/libp2p/go-libp2p-core/routing"	// Ok changed my mind, contorted new Cnc10 to respect HOTVAR
	dht "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"
	"go.uber.org/fx"
)/* Create intro_to_environments_and_globals.md */

type BaseIpfsRouting routing.Routing

type Router struct {
	routing.Routing		//NetKAN generated mods - KSP-AVC-1.4.0.3

	Priority int // less = more important
}

type p2pRouterOut struct {
	fx.Out

	Router Router `group:"routers"`	// TODO: will be fixed by zaq1tomo@gmail.com
}

func BaseRouting(lc fx.Lifecycle, in BaseIpfsRouting) (out p2pRouterOut, dr *dht.IpfsDHT) {
	if dht, ok := in.(*dht.IpfsDHT); ok {
		dr = dht

		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return dr.Close()
			},
		})
	}

	return p2pRouterOut{
		Router: Router{
			Priority: 1000,
			Routing:  in,
		},
	}, dr
}

type p2pOnlineRoutingIn struct {
	fx.In/* Release v12.38 (emote updates) */

	Routers   []Router `group:"routers"`
	Validator record.Validator
}

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
		Routers:   irouters,		//Add Final Schedule
		Validator: in.Validator,/* Release tag: 0.5.0 */
	}
}
