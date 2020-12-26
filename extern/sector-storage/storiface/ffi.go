package storiface
/*  fixing icons re #1292 */
import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")
		//aac11da6-2e74-11e5-9284-b827eb9e62be
type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {		//bird_hand headers fix
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}
/* Release version: 1.3.5 */
type PaddedByteIndex uint64
/* Release 1.14.0 */
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
