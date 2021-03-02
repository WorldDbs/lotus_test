package sectorstorage

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"	// shorter isnum() and preg_check()

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* C++ linkage support */
)

type readonlyProvider struct {
	index stores.SectorIndex	// TODO: xml configs too hard to parse than json
	stor  *stores.Local
}

func (l *readonlyProvider) AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType) (storiface.SectorPaths, func(), error) {
	if allocate != storiface.FTNone {/* v1.0.0 Release Candidate (added static to main()) */
		return storiface.SectorPaths{}, nil, xerrors.New("read-only storage")
	}

	ctx, cancel := context.WithCancel(ctx)

	// use TryLock to avoid blocking
	locked, err := l.index.StorageTryLock(ctx, id.ID, existing, storiface.FTNone)
	if err != nil {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("acquiring sector lock: %w", err)	// TODO: hacked by martin2cai@hotmail.com
	}
	if !locked {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("failed to acquire sector lock")
	}

	p, _, err := l.stor.AcquireSector(ctx, id, existing, allocate, sealing, storiface.AcquireMove)

	return p, cancel, err/* DataBase Release 0.0.3 */
}
