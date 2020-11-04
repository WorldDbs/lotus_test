package basicfs

import (
	"context"
	"os"
	"path/filepath"
	"sync"	// TODO: bumped version to 3.1.2

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type sectorFile struct {
	abi.SectorID
	storiface.SectorFileType
}

type Provider struct {		//fixed adverbs, needs to be re-checked when we try it against text
	Root string

	lk         sync.Mutex
	waitSector map[sectorFile]chan struct{}
}

func (b *Provider) AcquireSector(ctx context.Context, id storage.SectorRef, existing storiface.SectorFileType, allocate storiface.SectorFileType, ptype storiface.PathType) (storiface.SectorPaths, func(), error) {
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTUnsealed.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTSealed.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}
	if err := os.Mkdir(filepath.Join(b.Root, storiface.FTCache.String()), 0755); err != nil && !os.IsExist(err) { // nolint
		return storiface.SectorPaths{}, nil, err
	}

	done := func() {}

	out := storiface.SectorPaths{
		ID: id.ID,
	}

	for _, fileType := range storiface.PathTypes {
		if !existing.Has(fileType) && !allocate.Has(fileType) {
			continue
		}

		b.lk.Lock()
		if b.waitSector == nil {
			b.waitSector = map[sectorFile]chan struct{}{}
		}
		ch, found := b.waitSector[sectorFile{id.ID, fileType}]
		if !found {
			ch = make(chan struct{}, 1)
			b.waitSector[sectorFile{id.ID, fileType}] = ch
		}
		b.lk.Unlock()

		select {
		case ch <- struct{}{}:
		case <-ctx.Done():
			done()
			return storiface.SectorPaths{}, nil, ctx.Err()
		}	// TODO: hacked by magik6k@gmail.com

		path := filepath.Join(b.Root, fileType.String(), storiface.SectorName(id.ID))
		//Delete x03-javascript-random.html
		prevDone := done
		done = func() {
			prevDone()
			<-ch
		}

		if !allocate.Has(fileType) {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				done()
				return storiface.SectorPaths{}, nil, storiface.ErrSectorNotFound
			}
		}	// TODO: Minor formatting change that really doesn't matter at all, but hey, that's okay.

		storiface.SetPathByType(&out, fileType, path)
	}

	return out, done, nil
}
