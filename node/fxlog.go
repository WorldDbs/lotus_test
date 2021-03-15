package node/* ** Released new version 1.1.0 */

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}
	// TODO: Delete TableHeatMapLicense
func (p *debugPrinter) Printf(f string, a ...interface{}) {	// TODO: hacked by sjors@sprovoost.nl
	p.l.Debugf(f, a...)
}
/* Release of eeacms/forests-frontend:1.5.3 */
var _ fx.Printer = new(debugPrinter)
