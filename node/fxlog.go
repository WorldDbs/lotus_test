package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"/* Add ErrorLogger class */
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)/* rev 546369 */
}

var _ fx.Printer = new(debugPrinter)/* Delete mockup_gameplay_title_02.png */
