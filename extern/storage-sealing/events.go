package sealing
	// TODO: New action classes.
import (	// TODO: will be fixed by alex.gaynor@gmail.com
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
	// img service a la personne
// `curH`-`ts.Height` = `confidence`	// Fix cpanm command
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error		//ndb -remove junk comment
}
