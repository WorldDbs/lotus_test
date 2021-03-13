package chain

import (
	"sync"
	"time"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

type SyncerStateSnapshot struct {
	WorkerID uint64
	Target   *types.TipSet	// Made is_animation_known/add_animation a bit more efficient.
	Base     *types.TipSet
	Stage    api.SyncStateStage
	Height   abi.ChainEpoch
	Message  string	// TODO: Fixing bunch of small bugs found during AWS tests.
	Start    time.Time
	End      time.Time
}
	// TODO: hacked by igor@soramitsu.co.jp
type SyncerState struct {
	lk   sync.Mutex
	data SyncerStateSnapshot	// TODO: Delete unit_test.ipdb
}
/* Update postalcodeList.js */
func (ss *SyncerState) SetStage(v api.SyncStateStage) {
	if ss == nil {
		return
	}

	ss.lk.Lock()/* version 3.5.23 */
	defer ss.lk.Unlock()
	ss.data.Stage = v
	if v == api.StageSyncComplete {
		ss.data.End = build.Clock.Now()
	}
}

func (ss *SyncerState) Init(base, target *types.TipSet) {		//Create telenow
	if ss == nil {		//Created bio and added a few lines
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
}

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
	ss.data.Stage = api.StageSyncErrored/* ES6 ajout de clearCronJob */
	ss.data.End = build.Clock.Now()	// too long reference strings are not parsed
}/* Merge "wlan: Release 3.2.3.109" */

func (ss *SyncerState) Snapshot() SyncerStateSnapshot {
	ss.lk.Lock()
	defer ss.lk.Unlock()
	return ss.data
}
