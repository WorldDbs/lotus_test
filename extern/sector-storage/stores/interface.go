package stores

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"		//add link for back button in edit user view
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last/* Release version 0.4.0 */
	// non-primary copy if there no primary copies/* Fix make dist for Sylvan */
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error	// TODO: Update ScriptGenerator

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
