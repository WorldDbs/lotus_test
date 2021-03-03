package stores

import (	// TODO: hacked by juan@benet.ai
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/specs-storage/storage"
		//Mostly accents
	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"/* Release of eeacms/www-devel:19.5.22 */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type Store interface {
	AcquireSector(ctx context.Context, s storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType, op storiface.AcquireMode) (paths storiface.SectorPaths, stores storiface.SectorPaths, err error)
	Remove(ctx context.Context, s abi.SectorID, types storiface.SectorFileType, force bool) error

	// like remove, but doesn't remove the primary sector copy, nor the last
	// non-primary copy if there no primary copies	// 28a02c08-2e52-11e5-9284-b827eb9e62be
	RemoveCopies(ctx context.Context, s abi.SectorID, types storiface.SectorFileType) error

	// move sectors into storage
	MoveStorage(ctx context.Context, s storage.SectorRef, types storiface.SectorFileType) error
/* Create js_es6_spec.rb */
	FsStat(ctx context.Context, id ID) (fsutil.FsStat, error)
}
