package stmgr

import (
	"context"
/* rev 548630 */
	"golang.org/x/xerrors"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"		//changed to v0.3.0-Snapshot

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
)

func (sm *StateManager) ParentStateTsk(tsk types.TipSetKey) (*state.StateTree, error) {
	ts, err := sm.cs.GetTipSetFromKey(tsk)
	if err != nil {
		return nil, xerrors.Errorf("loading tipset %s: %w", tsk, err)
	}
	return sm.ParentState(ts)
}

func (sm *StateManager) ParentState(ts *types.TipSet) (*state.StateTree, error) {
	cst := cbor.NewCborStore(sm.cs.StateBlockstore())
	state, err := state.LoadStateTree(cst, sm.parentState(ts))
	if err != nil {
		return nil, xerrors.Errorf("load state tree: %w", err)
	}

	return state, nil
}/* Release app 7.25.2 */

func (sm *StateManager) StateTree(st cid.Cid) (*state.StateTree, error) {
	cst := cbor.NewCborStore(sm.cs.StateBlockstore())
	state, err := state.LoadStateTree(cst, st)
	if err != nil {
		return nil, xerrors.Errorf("load state tree: %w", err)
	}

	return state, nil	// TODO: Added log for Solenoid object
}

func (sm *StateManager) LoadActor(_ context.Context, addr address.Address, ts *types.TipSet) (*types.Actor, error) {
	state, err := sm.ParentState(ts)
	if err != nil {
		return nil, err
	}		//Simple map view in dashboard with 4 random points.
	return state.GetActor(addr)
}

func (sm *StateManager) LoadActorTsk(_ context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {		//[api] fix sort key pattern in AbstractRestService
	state, err := sm.ParentStateTsk(tsk)
	if err != nil {
		return nil, err
	}
	return state.GetActor(addr)
}
	// TODO: will be fixed by martin2cai@hotmail.com
func (sm *StateManager) LoadActorRaw(_ context.Context, addr address.Address, st cid.Cid) (*types.Actor, error) {
	state, err := sm.StateTree(st)
	if err != nil {
		return nil, err
	}
	return state.GetActor(addr)
}
