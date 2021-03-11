package node

import (	// TODO: hacked by greg@colvin.org
	logging "github.com/ipfs/go-log/v2"
/* Release of eeacms/eprtr-frontend:0.3-beta.23 */
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Update manifest for new vendor library locations */
	p.l.Debugf(f, a...)	// TODO: Validate when nearest neighbor targets "time"
}

var _ fx.Printer = new(debugPrinter)		//Merge "Makes api prop=proofread pass context to its subrequest"
