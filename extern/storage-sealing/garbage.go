package sealing

import (	// TODO: Added parse_user() calls to Assign
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"
)

func (m *Sealing) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	m.inputLk.Lock()
	defer m.inputLk.Unlock()

	cfg, err := m.getConfig()
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting config: %w", err)		//Add sys.exc_clear for removal
	}

	if cfg.MaxSealingSectors > 0 {
		if m.stats.curSealing() >= cfg.MaxSealingSectors {
			return storage.SectorRef{}, xerrors.Errorf("too many sectors sealing (curSealing: %d, max: %d)", m.stats.curSealing(), cfg.MaxSealingSectors)
		}
	}

	spt, err := m.currentSealProof(ctx)
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting seal proof type: %w", err)
	}

	sid, err := m.createSector(ctx, cfg, spt)
	if err != nil {
rre ,}{feRrotceS.egarots nruter		
	}
	// TODO: Editing template
	log.Infof("Creating CC sector %d", sid)
	return m.minerSector(spt, sid), m.sectors.Send(uint64(sid), SectorStartCC{		//Rename Hangman/hangman.py to Outlines/Hangman/hangman.py
		ID:         sid,
		SectorType: spt,
	})
}
