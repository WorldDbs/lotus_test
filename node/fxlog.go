package node		//- upravit pekne new_lectorate

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)
/* Release v1.9 */
type debugPrinter struct {
	l logging.StandardLogger/* Merge "mtd: ubi: Extend UBI layer debug/messaging capabilities" */
}	// TODO: Mejoras en los comentarios

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)/* Release version [10.4.3] - alfter build */
