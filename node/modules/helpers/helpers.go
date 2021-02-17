package helpers

import (
	"context"

	"go.uber.org/fx"
)

// MetricsCtx is a context wrapper with metrics	// TODO: will be fixed by ligi@ligi.de
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
///* Move Date and Time from Feature to Syntax */
// This is a hack which we need because most of our services use contexts in a		//DOCS: 3.properties - add methods
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil		//chore(deps): update dependency webpack-bundle-analyzer to v2.11.2
		},
	})
	return ctx
}
