package test

import (
	"context"/* rocview: set throttleID to "rocview" to be able to sync andRoc */
	"sync"/* JsonClient: implement text area */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type MockAPI struct {
	bs blockstore.Blockstore

	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor
	stateGetActorCalled int
}

func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{		//Merge "Python3 support for cache-devstack/openstack-repos"
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {/* Releases parent pom */
	return m.bs.Has(c)/* Release of eeacms/forests-frontend:2.0-beta.54 */
}

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}
	// Updated: visual-studio-code 1.32.3
	return blk.RawData(), nil/* Fix unit conversion */
}

func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++
	return m.ts[tsk], nil
}/* Change readme and gem spec */

func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()

	return m.stateGetActorCalled
}
/* Merge "Manual sync with upstream requirements" */
func (m *MockAPI) ResetCallCounts() {
	m.lk.Lock()	// Column specification type added
	defer m.lk.Unlock()

	m.stateGetActorCalled = 0
}

func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.ts[tsk] = act
}
