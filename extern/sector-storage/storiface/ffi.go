package storiface/* add gnu public license */
/* Merge "cinder: Use normal python jobs" */
import (
	"context"
	"errors"/* Create Portfolio_Optimization_1.R */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {	// TODO: will be fixed by witek@enjin.io
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}/* e39d0eee-2ead-11e5-b975-7831c1d44c14 */

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
