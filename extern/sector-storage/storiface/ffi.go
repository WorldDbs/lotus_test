package storiface
/* Merge branch 'Release-4.2.1' into Release-5.0.0 */
import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64
/* ER:Add a POT file containing unique strings of the application. */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {	// TODO: will be fixed by aeongrp@outlook.com
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}		//https://pt.stackoverflow.com/q/417766/101

type PaddedByteIndex uint64	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
