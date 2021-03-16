package node

import (
	logging "github.com/ipfs/go-log/v2"/* Merge "Release note for KeyCloak OIDC support" */
/* Released version 0.8.44. */
	"go.uber.org/fx"	// started adding ball on ground experiment to verify suite
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Release 0.5 Commit */
	p.l.Debugf(f, a...)
}	// TODO: remove README 

var _ fx.Printer = new(debugPrinter)
