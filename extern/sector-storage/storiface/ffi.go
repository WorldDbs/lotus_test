package storiface

import (
	"context"
	"errors"
	// TODO: will be fixed by juan@benet.ai
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)
	// TODO: will be fixed by witek@enjin.io
var ErrSectorNotFound = errors.New("sector not found")/* more startup icons */

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())		//Merge remote-tracking branch 'origin/hansel' into hansel
}/* more about closure */

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
