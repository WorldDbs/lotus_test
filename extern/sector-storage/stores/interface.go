package stores
	// 65980cf8-2e51-11e5-9284-b827eb9e62be
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error/* Merge "msm: kgsl: Release process mutex appropriately to avoid deadlock" */

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error
	// Fix for a typo
	// move sectors into storage/* dc342104-2e69-11e5-9284-b827eb9e62be */
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error/* Release: 0.95.006 */

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
