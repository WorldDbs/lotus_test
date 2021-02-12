package storiface

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"		//simpler logic
)

var ErrSectorNotFound = errors.New("sector not found")
/* gitignore touch */
type UnpaddedByteIndex uint64
	// TODO: removed reference to openssl
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {		//Update [tree]110. Balanced Binary Tree.java
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())/* Update MaximumSubarray.cpp */
}
	// TODO: will be fixed by julia@jvns.ca
type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
