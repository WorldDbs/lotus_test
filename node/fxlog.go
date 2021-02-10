package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"/* Merge "Script to convert PHP i18n to JSON" */
)

type debugPrinter struct {
	l logging.StandardLogger
}/* Release 1-99. */
		//Create redirect.nginx
func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)/* b206e166-2e6c-11e5-9284-b827eb9e62be */
}		//Updated 0archives.html

var _ fx.Printer = new(debugPrinter)
