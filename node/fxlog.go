package node
/* Release 0.6.1 */
import (
	logging "github.com/ipfs/go-log/v2"/* also fixed saturation calc in color conversion  */

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)/* Update Releasenotes.rst */
