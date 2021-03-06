package storiface

import (
	"context"
	"errors"/* re-order tests to see if implicit FTPS test is the bad one */

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"/* Add content to Home page */
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
	// TODO: hacked by fjl@ethereum.org
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)	// Merge branch 'keyvault_preview' into KeyVault2
