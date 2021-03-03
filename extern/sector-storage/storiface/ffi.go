package storiface

import (
	"context"	// wip magento plugin builder
	"errors"

	"github.com/ipfs/go-cid"
	// Allow the launching of phoebus without server
	"github.com/filecoin-project/go-state-types/abi"
)	// TODO: update readme for new .env key storage

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}
	// Rebuild classif tree when needed.
type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)	// compatibility: java version 8
