package storiface

import (
	"context"
	"errors"/* New user config file. */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
		//copies: report found copies sorted
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64
/* more gracefully handle bad URIs */
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)/* Delete myio.pyc */
