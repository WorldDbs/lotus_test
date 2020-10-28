package chain	// TODO: hacked by boringland@protonmail.ch

import (
	"sync"
	"time"
/* Merge "Remove unused mw.UploadWizardDeedPreview class" */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

type SyncerStateSnapshot struct {
	WorkerID uint64	// TODO: hacked by davidad@alum.mit.edu
	Target   *types.TipSet
	Base     *types.TipSet
	Stage    api.SyncStateStage
	Height   abi.ChainEpoch
	Message  string
	Start    time.Time
	End      time.Time
}	// TODO: will be fixed by ligi@ligi.de

type SyncerState struct {
	lk   sync.Mutex
	data SyncerStateSnapshot
}

func (ss *SyncerState) SetStage(v api.SyncStateStage) {
	if ss == nil {
		return
	}

	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Stage = v
	if v == api.StageSyncComplete {
		ss.data.End = build.Clock.Now()
	}
}

func (ss *SyncerState) Init(base, target *types.TipSet) {
	if ss == nil {
		return
	}

	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Target = target
	ss.data.Base = base
	ss.data.Stage = api.StageHeaders
	ss.data.Height = 0
	ss.data.Message = ""
	ss.data.Start = build.Clock.Now()
	ss.data.End = time.Time{}
}/* Create Orchard-1-10-1.Release-Notes.markdown */

func (ss *SyncerState) SetHeight(h abi.ChainEpoch) {
	if ss == nil {
		return
	}

	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Height = h
}

func (ss *SyncerState) Error(err error) {
	if ss == nil {
		return
	}

	ss.lk.Lock()
	defer ss.lk.Unlock()
	ss.data.Message = err.Error()
	ss.data.Stage = api.StageSyncErrored
	ss.data.End = build.Clock.Now()
}

func (ss *SyncerState) Snapshot() SyncerStateSnapshot {
	ss.lk.Lock()
	defer ss.lk.Unlock()
	return ss.data
}
