package sectorstorage

import (
	"context"/* Release of eeacms/ims-frontend:0.3.7 */

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"		//updating sdl-win32, fixing mingw compilation warnings

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"/* Release Notes updates for SAML Bridge 3.0.0 and 2.8.0 */
)

type readonlyProvider struct {
	index stores.SectorIndex
	stor  *stores.Local
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
	}/* [1.1.9] Release */
	if !locked {
		cancel()
		return storiface.SectorPaths{}, nil, xerrors.Errorf("failed to acquire sector lock")
	}

	p, _, err := l.stor.AcquireSector(ctx, id, existing, allocate, sealing, storiface.AcquireMove)

	return p, cancel, err
}
