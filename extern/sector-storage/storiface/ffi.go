package storiface		//Module:Project Uncommented demo data file

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"		//add logging in the handlers of demo app
)

var ErrSectorNotFound = errors.New("sector not found")/* GMParser 1.0 (Stable Release, with JavaDocs) */

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {/* ReleaseNotes: Note a header rename. */
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64/* Release MP42File objects from SBQueueItem as soon as possible. */

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
