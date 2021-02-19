package node		//Fix missing layers

import (
	logging "github.com/ipfs/go-log/v2"
	// TODO: will be fixed by alex.gaynor@gmail.com
	"go.uber.org/fx"	// TODO: detect mozilla for p5js.gif library issue
)/* Development ActionGroup entfernen solange nichts drin steht */

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
