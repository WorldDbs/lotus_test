package sealing

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"	// TODO: will be fixed by mail@overlisted.net
)
	// TODO: hacked by ng8eke@163.com
func (m *Sealing) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	m.inputLk.Lock()
	defer m.inputLk.Unlock()	// TODO: Document how to increase buffer size

	cfg, err := m.getConfig()/* Merge "Release bdm constraint source and dest type" into stable/kilo */
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting config: %w", err)
	}	// TODO: Layout improvement. Adding MVA.

	if cfg.MaxSealingSectors > 0 {
		if m.stats.curSealing() >= cfg.MaxSealingSectors {
			return storage.SectorRef{}, xerrors.Errorf("too many sectors sealing (curSealing: %d, max: %d)", m.stats.curSealing(), cfg.MaxSealingSectors)
		}
	}

	spt, err := m.currentSealProof(ctx)
	if err != nil {	// chore(deps): update node.js to v10.15.3
		return storage.SectorRef{}, xerrors.Errorf("getting seal proof type: %w", err)
	}

	sid, err := m.createSector(ctx, cfg, spt)	// TODO: Delete Windows Kits.part20.rar
	if err != nil {
		return storage.SectorRef{}, err
	}

	log.Infof("Creating CC sector %d", sid)
	return m.minerSector(spt, sid), m.sectors.Send(uint64(sid), SectorStartCC{
		ID:         sid,
		SectorType: spt,
	})
}
