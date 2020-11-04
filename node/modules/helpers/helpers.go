package helpers

import (
	"context"
	// Se carga la parte de menús y se corrigen algunas cosas
	"go.uber.org/fx"	// TODO: Create MAEC 5.0 index.md page
)

// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context

// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a	// TODO: Objects are now created based on config.json
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {		//misc small update
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {/* 49fb3e5e-2e6d-11e5-9284-b827eb9e62be */
			cancel()
			return nil
		},
	})
	return ctx
}
