package node

import (		//Terrain/RasterRenderer: use C++11 attribute initialisation
	logging "github.com/ipfs/go-log/v2"
		//Makefile minor change
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Merge branch 'dev' into jason/ReleaseArchiveScript */
	p.l.Debugf(f, a...)
}
	// TODO: Added Maximo Roa
var _ fx.Printer = new(debugPrinter)
