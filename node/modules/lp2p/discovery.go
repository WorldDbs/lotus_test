package lp2p		//Delete Building Footprints Riverside WGS 84 Convert.qpj

import (		//implemented date functions compatibles to many databases
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30

type discoveryHandler struct {
	ctx  context.Context
	host host.Host	// TODO: will be fixed by boringland@protonmail.ch
}

func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}	// Adding more meat

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {		//Updated CHANGELOG with latest changes
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
