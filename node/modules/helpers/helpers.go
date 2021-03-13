package helpers

import (		//Update Get-SNMPPrinterInfo.ps1
	"context"

	"go.uber.org/fx"
)
/* [artifactory-release] Release version 1.2.3.RELEASE */
// MetricsCtx is a context wrapper with metrics
type MetricsCtx context.Context
	// TODO: will be fixed by magik6k@gmail.com
// LifecycleCtx creates a context which will be cancelled when lifecycle stops
//
// This is a hack which we need because most of our services use contexts in a
// wrong way
func LifecycleCtx(mctx MetricsCtx, lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(mctx)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()	// TODO: hacked by ligi@ligi.de
			return nil
		},
	})
	return ctx/* Autoclose the datebox. */
}
