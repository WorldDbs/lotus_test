package helpers
/* Add missing word in PreRelease.tid */
import (
	"context"

	"go.uber.org/fx"
)
/* Release CAPO 0.3.0-rc.0 image */
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops/* Use masking instead of EOW padding */
//
// This is a hack which we need because most of our services use contexts in a/* Update hotel_create.html */
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {/* Enforce ordering. */
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{/* Merge "input: touchpanel: Release all touches during suspend" */
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})
	return ctx
}
