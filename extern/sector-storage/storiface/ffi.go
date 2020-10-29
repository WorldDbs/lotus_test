package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"
/* Version Release (Version 1.6) */
	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
/* Added 'the most important changes since 0.6.1' in Release_notes.txt */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
