package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"/* Release 2.4.0 */
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	"github.com/filecoin-project/go-state-types/abi"
)
/* Delete bold.gif */
var ErrSectorNotFound = errors.New("sector not found")		//Added build configuration topic in Development Environment

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
