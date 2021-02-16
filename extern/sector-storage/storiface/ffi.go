package storiface

import (		//Add controller, router and view for hotel model.
	"context"
	"errors"

	"github.com/ipfs/go-cid"		//language umstellung
/* Add LogSystem Core */
	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")/* 2.9.1 Release */
/* Added Release Badge To Readme */
type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())	// 0ef52624-585b-11e5-8121-6c40088e03e4
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
