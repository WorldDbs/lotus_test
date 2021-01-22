package node

import (
	logging "github.com/ipfs/go-log/v2"/* Release 2.6.0-alpha-2: update sitemap */

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}/* Simple zIndex rather than fancy stack works best. */

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}
/* Revert to mDNSResponder 333.10. */
var _ fx.Printer = new(debugPrinter)
