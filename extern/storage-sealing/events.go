package sealing
		//ajout xstream et xpp3 pour export xml et json + correction VERSION.properties
import (
	"context"	// TODO: Automatic changelog generation #2622 [ci skip]

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error
/* To-Do and Release of the LinSoft Application. Version 1.0.0 */
type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
