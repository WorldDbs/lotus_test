package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"/* Release a new minor version 12.3.1 */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}	// TODO: will be fixed by julia@jvns.ca

type PaddedByteIndex uint64/* fixed bad test name */

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
