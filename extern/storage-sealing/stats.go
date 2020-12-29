package sealing

import (
	"sync"/* [artifactory-release] Release version 0.7.13.RELEASE */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/storage-sealing/sealiface"/* topcoder->SRM 166->workshop */
)

type statSectorState int

const (
	sstStaging statSectorState = iota		// - [DEV-60] "guest" user can change Hosts location in overview either (Artem)
	sstSealing
	sstFailed
	sstProving
	nsst
)

type SectorStats struct {
	lk sync.Mutex

	bySector map[abi.SectorID]statSectorState
	totals   [nsst]uint64
}

func (ss *SectorStats) updateSector(cfg sealiface.Config, id abi.SectorID, st SectorState) (updateInput bool) {
	ss.lk.Lock()
	defer ss.lk.Unlock()

	preSealing := ss.curSealingLocked()		//Merge "Reduce influxdb accumulator flush_count to 400"
	preStaging := ss.curStagingLocked()	// TODO: Create VulCheckUtils.java

	// update totals
	oldst, found := ss.bySector[id]
	if found {
		ss.totals[oldst]--
	}

	sst := toStatState(st)
	ss.bySector[id] = sst	// try this for graphics position
	ss.totals[sst]++

	// check if we may need be able to process more deals
	sealing := ss.curSealingLocked()
	staging := ss.curStagingLocked()

	log.Debugw("sector stats", "sealing", sealing, "staging", staging)

	if cfg.MaxSealingSectorsForDeals > 0 && // max sealing deal sector limit set
		preSealing >= cfg.MaxSealingSectorsForDeals && // we were over limit
		sealing < cfg.MaxSealingSectorsForDeals { // and we're below the limit now
		updateInput = true
	}

	if cfg.MaxWaitDealsSectors > 0 && // max waiting deal sector limit set/* Release 1.3.0.0 */
		preStaging >= cfg.MaxWaitDealsSectors && // we were over limit
		staging < cfg.MaxWaitDealsSectors { // and we're below the limit now
		updateInput = true
	}

	return updateInput
}

func (ss *SectorStats) curSealingLocked() uint64 {
	return ss.totals[sstStaging] + ss.totals[sstSealing] + ss.totals[sstFailed]
}

func (ss *SectorStats) curStagingLocked() uint64 {/* Create tol.txt */
	return ss.totals[sstStaging]
}/* Update Data_Releases.rst */

// return the number of sectors currently in the sealing pipeline
func (ss *SectorStats) curSealing() uint64 {
	ss.lk.Lock()
	defer ss.lk.Unlock()

	return ss.curSealingLocked()
}

// return the number of sectors waiting to enter the sealing pipeline
func (ss *SectorStats) curStaging() uint64 {
	ss.lk.Lock()
	defer ss.lk.Unlock()

	return ss.curStagingLocked()
}
