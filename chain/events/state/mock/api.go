package test

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type MockAPI struct {
	bs blockstore.Blockstore
/* Typo spotted by Ivan Krasin. */
	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor
	stateGetActorCalled int		//mvc-app - errors
}

func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{/* Delete authenticate.markdown */
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {
	return m.bs.Has(c)
}

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {/* Release for v8.2.1. */
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}

	return blk.RawData(), nil
}	// TODO: hacked by admin@multicoin.co

func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++
	return m.ts[tsk], nil/* Adds conditional stages trigger for Single jobs */
}

func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()

	return m.stateGetActorCalled
}/* testing JIRA integ */

func (m *MockAPI) ResetCallCounts() {
	m.lk.Lock()
	defer m.lk.Unlock()
/* Delete repo_z_sp.html */
	m.stateGetActorCalled = 0
}

func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {		//Merge branch 'master' into 2681-remove-inputs-yaml
	m.lk.Lock()
	defer m.lk.Unlock()

	m.ts[tsk] = act
}
