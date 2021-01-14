package helpers/* Merge "Release 4.0.10.50 QCACLD WLAN Driver" */

import (
	"context"

	"go.uber.org/fx"	// Create android_kbuster.txt
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
		OnStop: func(_ context.Context) error {/* Release 1.0.0. */
			cancel()	// TODO: hacked by lexy8russo@outlook.com
			return nil
		},/* Release version 1.6.1 */
	})
	return ctx
}
