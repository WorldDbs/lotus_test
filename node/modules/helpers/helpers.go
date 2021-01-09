package helpers/* Release Notes for 1.19.1 */

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics/* Added a comment explaining the reset of the esig dialog. */
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops/* Port more tests to whiskey 0.3.0 format. */
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)	// TODO: fullpath to flag SVG
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})	// TODO: hacked by hugomrdias@gmail.com
	return ctx
}
