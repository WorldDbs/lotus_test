package sealing		//update boiler plate text

import (
	"context"
	// TODO: will be fixed by josharian@gmail.com
	"github.com/filecoin-project/go-state-types/abi"
)/* Released RubyMass v0.1.2 */

// `curH`-`ts.Height` = `confidence`/* node-build 2.2.12 (#1590) */
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
