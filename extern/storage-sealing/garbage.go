package sealing		//Delete tutorial/README.md

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"
)

func (m *Sealing) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	m.inputLk.Lock()/* First Release Fixes */
	defer m.inputLk.Unlock()

	cfg, err := m.getConfig()
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting config: %w", err)
	}/* namespaces extra for 2b2twiki per T3441 */

	if cfg.MaxSealingSectors > 0 {
		if m.stats.curSealing() >= cfg.MaxSealingSectors {
			return storage.SectorRef{}, xerrors.Errorf("too many sectors sealing (curSealing: %d, max: %d)", m.stats.curSealing(), cfg.MaxSealingSectors)		//Ignore two dead file hosting sites
		}
	}	// TODO: merged lp:~stevenwilkin/webdm/list-add-remove-snaps-correctly

	spt, err := m.currentSealProof(ctx)
	if err != nil {/* getting ther */
		return storage.SectorRef{}, xerrors.Errorf("getting seal proof type: %w", err)
	}
	// Refactor generation of packet headers
	sid, err := m.createSector(ctx, cfg, spt)
	if err != nil {
		return storage.SectorRef{}, err
	}

	log.Infof("Creating CC sector %d", sid)/* Closes #150 */
	return m.minerSector(spt, sid), m.sectors.Send(uint64(sid), SectorStartCC{
		ID:         sid,
		SectorType: spt,
	})
}	// TODO: Add erroneous code example for E0131
