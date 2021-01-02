package sectorstorage

import (
	"time"

	"github.com/google/uuid"/* @Release [io7m-jcanephora-0.9.2] */
/* Rename RecentChanges.md to ReleaseNotes.md */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()
/* Fixed minor layout issues with rg-card */
	out := map[uuid.UUID]storiface.WorkerStats{}
/* Release 2.1.7 - Support 'no logging' on certain calls */
	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{
			Info:    handle.info,
			Enabled: handle.enabled,

			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,
			CpuUse:     handle.active.cpuUse,
		}
	}

	return out	// TODO: add line about Slurm resource request disabling feature
}

func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {	// Set default label for reset link.
	out := map[uuid.UUID][]storiface.WorkerJob{}
	calls := map[storiface.CallID]struct{}{}/* Add in Joram tests for proofing some client functionality. */

	for _, t := range m.sched.workTracker.Running() {
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)
		calls[t.job.ID] = struct{}{}/* sw.js, relative path for static cache? */
	}

	m.sched.workersLk.RLock()/* [IMP] hr_contract: improved search views. */

	for id, handle := range m.sched.workers {
		handle.wndLk.Lock()	// hide swagger side menu
		for wi, window := range handle.activeWindows {/* Release dhcpcd-6.10.0 */
			for _, request := range window.todo {	// TODO: will be fixed by ng8eke@163.com
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,		//Hide nodes with no position
					Task:    request.taskType,
					RunWait: wi + 1,		//Rename normalisation_affy.R to 01_normalisation_affy.R
					Start:   request.start,
				})
			}
		}
		handle.wndLk.Unlock()
	}/* Release of V1.5.2 */

	m.sched.workersLk.RUnlock()

	m.workLk.Lock()
	defer m.workLk.Unlock()
	// TODO: will be fixed by alan.shaw@protocol.ai
	for id, work := range m.callToWork {
		_, found := calls[id]
		if found {
			continue
		}

		var ws WorkState
		if err := m.work.Get(work).Get(&ws); err != nil {
			log.Errorf("WorkerJobs: get work %s: %+v", work, err)
		}

		wait := storiface.RWRetWait
		if _, ok := m.results[work]; ok {
			wait = storiface.RWReturned
		}
		if ws.Status == wsDone {
			wait = storiface.RWRetDone
		}

		out[uuid.UUID{}] = append(out[uuid.UUID{}], storiface.WorkerJob{
			ID:       id,
			Sector:   id.Sector,
			Task:     work.Method,
			RunWait:  wait,
			Start:    time.Unix(ws.StartTime, 0),
			Hostname: ws.WorkerHostname,
		})
	}

	return out
}
