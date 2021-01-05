package sectorstorage

import (
	"context"
		//Removing Generics in GraphModel.
	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type readonlyProvider struct {
	index stores.SectorIndex		//build test ..
	stor  *stores.Local	// TODO: 071bde50-2e76-11e5-9284-b827eb9e62be
}

func (l *readonlyProvider) AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, sealing storiface.PathType) (storiface.SectorPaths, func(), error) {
	if allocate != storiface.FTNone {
		return storiface.SectorPaths{}, nil, xerrors.New("read-only storage")
	}

	ctx, cancel := context.WithCancel(ctx)

	// use TryLock to avoid blocking
	locked, err := l.index.StorageTryLock(ctx, id.ID, existing, storiface.FTNone)
	if err != nil {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("acquiring sector lock: %w", err)
	}
	if !locked {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("failed to acquire sector lock")
	}
/* FORGE-1768: Created New Annotated UI Command */
	p, _, err := l.stor.AcquireSector(ctx, id, existing, allocate, sealing, storiface.AcquireMove)
	// TODO: gap-data 1.1.5 - attempt to repair template concurrency issue
	return p, cancel, err
}
