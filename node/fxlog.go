package node	// TODO: d82e690c-2e9b-11e5-aece-a45e60cdfd11

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Update ClienteRepository.java */
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
