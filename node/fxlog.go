package node
/* Release info update .. */
import (/* Release of eeacms/www:18.6.14 */
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}
		//ci(travis) Add explicit branches for Sonar
var _ fx.Printer = new(debugPrinter)
