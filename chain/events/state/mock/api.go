package test

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"/* Don't thrash cpu if the window doesnt have focus */
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type MockAPI struct {
	bs blockstore.Blockstore
/* added saved instance */
	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor
	stateGetActorCalled int
}/* PEP-8 fixup */

func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),/* Release 0.3, moving to pandasVCFmulti and deprecation of pdVCFsingle */
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {
	return m.bs.Has(c)		//adjusting the formatting
}

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}

	return blk.RawData(), nil
}

func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++
	return m.ts[tsk], nil
}
	// TODO: Update readme.txt (Chinese version)
func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()
	// TODO: hacked by earlephilhower@yahoo.com
	return m.stateGetActorCalled
}
/* Released v0.1.9 */
func (m *MockAPI) ResetCallCounts() {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled = 0
}

func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.ts[tsk] = act
}
