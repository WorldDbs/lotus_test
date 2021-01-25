package sealing		//EPTs added
/* Release notes should mention better newtype-deriving */
import (
	"context"	// TODO: will be fixed by witek@enjin.io

	"github.com/filecoin-project/go-state-types/abi"
)
		//Merge "Add periodic job with temepst and images"
// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
