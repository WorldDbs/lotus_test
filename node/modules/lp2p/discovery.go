package lp2p

import (/* Delete GraphicsEngine.js */
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/fx"		//Delete Sites.js

	"github.com/filecoin-project/lotus/node/modules/helpers"
)

const discoveryConnTimeout = time.Second * 30	// Merge "xenapi: add username to vncviewer command"

type discoveryHandler struct {
	ctx  context.Context
	host host.Host
}
		//add freertos code
func (dh *discoveryHandler) HandlePeerFound(p peer.AddrInfo) {
	log.Warnw("discovred peer", "peer", p)
	ctx, cancel := context.WithTimeout(dh.ctx, discoveryConnTimeout)
	defer cancel()
	if err := dh.host.Connect(ctx, p); err != nil {/* Merge "wlan: Release 3.2.3.123" */
		log.Warnw("failed to connect to peer found by discovery", "error", err)
	}
}/* Release of eeacms/www:18.7.20 */

func DiscoveryHandler(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) *discoveryHandler {
	return &discoveryHandler{
		ctx:  helpers.LifecycleCtx(mctx, lc),
		host: host,
	}
}
