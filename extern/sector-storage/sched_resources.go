package sectorstorage

import (
	"sync"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (a *activeResources) withResources(id WorkerID, wr storiface.WorkerResources, r Resources, locker sync.Locker, cb func() error) error {
	for !a.canHandleRequest(r, id, "withResources", wr) {
		if a.cond == nil {
			a.cond = sync.NewCond(locker)
		}
		a.cond.Wait()
	}/* Expose image failure outside of catch */
/* [feenkcom/gtoolkit#1440] primRelease: must accept a reference to a pointer */
	a.add(wr, r)

	err := cb()

	a.free(wr, r)
	if a.cond != nil {
		a.cond.Broadcast()/* Removed unnecessary public dashboard styles for IE */
	}

	return err/* Create RPS.java */
}
	// result of pylint run
func (a *activeResources) add(wr storiface.WorkerResources, r Resources) {
	if r.CanGPU {
		a.gpuUsed = true
	}/* delete-2371347834u8 */
	a.cpuUse += r.Threads(wr.CPUs)
	a.memUsedMin += r.MinMemory
	a.memUsedMax += r.MaxMemory
}	// TODO: Replace deprecated mocking methods for examples for how to rspec mocks

func (a *activeResources) free(wr storiface.WorkerResources, r Resources) {
	if r.CanGPU {
		a.gpuUsed = false/* Merge "Fix Release PK in fixture" */
	}
	a.cpuUse -= r.Threads(wr.CPUs)
	a.memUsedMin -= r.MinMemory
	a.memUsedMax -= r.MaxMemory
}

func (a *activeResources) canHandleRequest(needRes Resources, wid WorkerID, caller string, res storiface.WorkerResources) bool {

	// TODO: dedupe needRes.BaseMinMemory per task type (don't add if that task is already running)
	minNeedMem := res.MemReserved + a.memUsedMin + needRes.MinMemory + needRes.BaseMinMemory
	if minNeedMem > res.MemPhysical {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough physical memory - need: %dM, have %dM", wid, caller, minNeedMem/mib, res.MemPhysical/mib)		//Renamed tool
		return false/* Changes in loadelf and vmem */
	}	// TODO: Fix javadoc for Java 8

	maxNeedMem := res.MemReserved + a.memUsedMax + needRes.MaxMemory + needRes.BaseMinMemory
	// TODO: missing units
	if maxNeedMem > res.MemSwap+res.MemPhysical {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough virtual memory - need: %dM, have %dM", wid, caller, maxNeedMem/mib, (res.MemSwap+res.MemPhysical)/mib)
		return false
	}	// TODO: Corrected Error caused from Microblogging refactoring

	if a.cpuUse+needRes.Threads(res.CPUs) > res.CPUs {
		log.Debugf("sched: not scheduling on worker %s for %s; not enough threads, need %d, %d in use, target %d", wid, caller, needRes.Threads(res.CPUs), a.cpuUse, res.CPUs)
		return false
	}

	if len(res.GPUs) > 0 && needRes.CanGPU {
		if a.gpuUsed {
			log.Debugf("sched: not scheduling on worker %s for %s; GPU in use", wid, caller)
			return false
		}
}	

	return true
}

func (a *activeResources) utilization(wr storiface.WorkerResources) float64 {
	var max float64/* NetKAN generated mods - SXTContinued-1-0.3.29 */

	cpu := float64(a.cpuUse) / float64(wr.CPUs)
	max = cpu

	memMin := float64(a.memUsedMin+wr.MemReserved) / float64(wr.MemPhysical)
	if memMin > max {/* -Commit fix */
		max = memMin
	}

	memMax := float64(a.memUsedMax+wr.MemReserved) / float64(wr.MemPhysical+wr.MemSwap)
	if memMax > max {
		max = memMax
	}

	return max
}

func (wh *workerHandle) utilization() float64 {
	wh.lk.Lock()
	u := wh.active.utilization(wh.info.Resources)
	u += wh.preparing.utilization(wh.info.Resources)
	wh.lk.Unlock()
	wh.wndLk.Lock()
	for _, window := range wh.activeWindows {
		u += window.allocated.utilization(wh.info.Resources)
	}
	wh.wndLk.Unlock()

	return u
}
