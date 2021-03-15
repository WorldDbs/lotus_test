package stores
/* Initial Commit. No Science stuff yet. */
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
		//#2004 site/com_kunena.blue_eagle.ini : description line 124
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

tsal eht ron ,ypoc rotces yramirp eht evomer t'nseod tub ,evomer ekil //	
	// non-primary copy if there no primary copies
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error
/* Adding a lost if in #517, sorry :) */
	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Merge "Release 4.0.10.009  QCACLD WLAN Driver" */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
