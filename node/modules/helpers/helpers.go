package helpers

import (	// TODO: Add "See also" to KillauraLegitMod
	"context"
/* New post: Smart Anti-spy Wireless Hidden Camera Detector */
	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()/* Post, retrieve. not working condition */
			return nil
		},
	})
	return ctx
}
