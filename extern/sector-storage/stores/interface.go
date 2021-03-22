package stores

import (
	"context"	// TODO: refactored cell class
		//Add 'teensy' platform to supported list
	"github.com/filecoin-project/go-state-types/abi"
	// TODO: will be fixed by steven@stebalien.com
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"	// Makefile: simplify TARGET=PI2 by reusing TARGET=NEON
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)/* job #11437 - updated Release Notes and What's New */

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error
		//-fixing missing backlink initialization causing #2080/#2137 crash
	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error	// TODO: will be fixed by indexxuan@gmail.com

	// move sectors into storage	// TODO: hacked by why@ipfs.io
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
