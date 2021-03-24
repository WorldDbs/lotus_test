package stores

import (
	"context"
	// TODO: Use Laravel database configuration by default
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error/* smaller scheduler batches */

	// move sectors into storage	// TODO: will be fixed by cory@protocol.ai
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
