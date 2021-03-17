package helpers	// TODO: hacked by why@ipfs.io

import (
	"context"/* Release as v0.10.1 */

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)	// Create technology_planning.py
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {/* feat: inline arrow function */
			cancel()
			return nil
		},
	})
	return ctx
}
