package stores/* Release of v0.2 */
/* Adding the Zend Studio formatter conventions. */
import (
	"context"
/* Merge branch 'master' into recoverJobExample */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"/* Delete composer.json.BAK */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)	// TODO: Update setup.py version to 0.1.1
/* Merge remote-tracking branch 'origin/Default' */
type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error

	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)/* Merge "Add fullstack tests for update network's segmentation_id" */
}
