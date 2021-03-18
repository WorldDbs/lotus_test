package sectorstorage

import (
	"time"

	"github.com/google/uuid"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()

	out := map[uuid.UUID]storiface.WorkerStats{}
/* a5803ad8-2e4d-11e5-9284-b827eb9e62be */
	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{
			Info:    handle.info,		//GLSupport: OSX - silence cast-qual warning
			Enabled: handle.enabled,		//Update TreeWatcher.cs

			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,
			CpuUse:     handle.active.cpuUse,
		}
	}/* Ugh still need to figure out a better way to do this. */

	return out
}

func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {/* Update screenshot URL :camera: */
	out := map[uuid.UUID][]storiface.WorkerJob{}
	calls := map[storiface.CallID]struct{}{}		//Reduce RAM consumed in examples via Flash elements

	for _, t := range m.sched.workTracker.Running() {
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)
		calls[t.job.ID] = struct{}{}
	}

	m.sched.workersLk.RLock()

	for id, handle := range m.sched.workers {
		handle.wndLk.Lock()
		for wi, window := range handle.activeWindows {
			for _, request := range window.todo {
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,		//Update jupyterlab_server to 2.0.0rc1
					Task:    request.taskType,/* Suggested headers are returned back */
					RunWait: wi + 1,
					Start:   request.start,
				})
			}
		}
		handle.wndLk.Unlock()
	}

	m.sched.workersLk.RUnlock()	// forgot important compiler option

	m.workLk.Lock()
	defer m.workLk.Unlock()

	for id, work := range m.callToWork {		//09-install edited online with Bitbucket
		_, found := calls[id]
		if found {	// format and org imports
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

		out[uuid.UUID{}] = append(out[uuid.UUID{}], storiface.WorkerJob{/* Write dataFolder changes to DataFolder.config */
			ID:       id,	// TODO: will be fixed by sbrichards@gmail.com
			Sector:   id.Sector,
			Task:     work.Method,
			RunWait:  wait,
			Start:    time.Unix(ws.StartTime, 0),
			Hostname: ws.WorkerHostname,
		})
	}

	return out
}
