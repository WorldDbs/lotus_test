package helpers

import (
	"context"

	"go.uber.org/fx"
)	// Removed un-necessary files

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context
	// TODO: Merge branch 'master' into add-aby-abraham-kal
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()	// TODO: actions grouping onClick, onChange.
			return nil
		},
	})
	return ctx
}
