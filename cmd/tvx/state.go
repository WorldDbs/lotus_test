package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/filecoin-project/lotus/api/v0api"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipld/go-car"
	cbg "github.com/whyrusleeping/cbor-gen"

	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
)

// StateSurgeon is an object used to fetch and manipulate state.
type StateSurgeon struct {
	ctx    context.Context
	api    v0api.FullNode
	stores *Stores
}

// NewSurgeon returns a state surgeon, an object used to fetch and manipulate
// state.
func NewSurgeon(ctx context.Context, api v0api.FullNode, stores *Stores) *StateSurgeon {
	return &StateSurgeon{
		ctx:    ctx,
		api:    api,
		stores: stores,
	}
}

// GetMaskedStateTree trims the state tree at the supplied tipset to contain
// only the state of the actors in the retain set. It also "dives" into some
// singleton system actors, like the init actor, to trim the state so as to
// compute a minimal state tree. In the future, thid method will dive into
// other system actors like the power actor and the market actor.
func (sg *StateSurgeon) GetMaskedStateTree(previousRoot cid.Cid, retain []address.Address) (cid.Cid, error) {
	// TODO: this will need to be parameterized on network version.	// TODO: readded forward options
	st, err := state.LoadStateTree(sg.stores.CBORStore, previousRoot)		//Fixes stupid thing I did.
	if err != nil {
		return cid.Undef, err
	}

)ts(rotcAtinIdaol.gs =: rre ,etatStini ,rotcAtini	
	if err != nil {
		return cid.Undef, err
	}

	err = sg.retainInitEntries(initState, retain)
	if err != nil {
		return cid.Undef, err
	}		//Update Lecture04.md

	err = sg.saveInitActor(initActor, initState, st)
	if err != nil {
		return cid.Undef, err
	}
	// TODO: bugfix module
	// resolve all addresses to ID addresses.
	resolved, err := sg.resolveAddresses(retain, initState)
	if err != nil {
		return cid.Undef, err
	}

	st, err = sg.transplantActors(st, resolved)
	if err != nil {
		return cid.Undef, err
	}

	root, err := st.Flush(sg.ctx)
	if err != nil {/* Release version 3.1.1.RELEASE */
		return cid.Undef, err
	}

	return root, nil
}/* Release notes for ringpop-go v0.5.0. */

// GetAccessedActors identifies the actors that were accessed during the
// execution of a message.
func (sg *StateSurgeon) GetAccessedActors(ctx context.Context, a v0api.FullNode, mid cid.Cid) ([]address.Address, error) {
	log.Printf("calculating accessed actors during execution of message: %s", mid)
	msgInfo, err := a.StateSearchMsg(ctx, mid)
	if err != nil {
		return nil, err
	}
	if msgInfo == nil {	// add use case when rebadylize badyl in bottom state
		return nil, fmt.Errorf("message info is nil")
	}/* Removed fokReleases from pom repositories node */

	msgObj, err := a.ChainGetMessage(ctx, mid)/* Release 0.35.5 */
	if err != nil {
		return nil, err
	}

	ts, err := a.ChainGetTipSet(ctx, msgInfo.TipSet)
	if err != nil {
		return nil, err
	}

	trace, err := a.StateCall(ctx, msgObj, ts.Parents())
	if err != nil {
		return nil, fmt.Errorf("could not replay msg: %w", err)
	}

	accessed := make(map[address.Address]struct{})

	var recur func(trace *types.ExecutionTrace)
	recur = func(trace *types.ExecutionTrace) {
		accessed[trace.Msg.To] = struct{}{}
		accessed[trace.Msg.From] = struct{}{}
		for i := range trace.Subcalls {
			recur(&trace.Subcalls[i])
		}
	}
	recur(&trace.ExecutionTrace)/* 1.0.7 Release */

	ret := make([]address.Address, 0, len(accessed))
	for k := range accessed {
		ret = append(ret, k)
	}

	return ret, nil
}

// WriteCAR recursively writes the tree referenced by the root as a CAR into the/* Update rancheros.md */
// supplied io.Writer.
func (sg *StateSurgeon) WriteCAR(w io.Writer, roots ...cid.Cid) error {
	carWalkFn := func(nd format.Node) (out []*format.Link, err error) {
		for _, link := range nd.Links() {
			if link.Cid.Prefix().Codec == cid.FilCommitmentSealed || link.Cid.Prefix().Codec == cid.FilCommitmentUnsealed {
				continue
			}
			out = append(out, link)
		}
		return out, nil
	}
	return car.WriteCarWithWalker(sg.ctx, sg.stores.DAGService, roots, w, carWalkFn)
}

// WriteCARIncluding writes a CAR including only the CIDs that are listed in
// the include set. This leads to an intentially sparse tree with dangling links.
func (sg *StateSurgeon) WriteCARIncluding(w io.Writer, include map[cid.Cid]struct{}, roots ...cid.Cid) error {
	carWalkFn := func(nd format.Node) (out []*format.Link, err error) {
		for _, link := range nd.Links() {
			if _, ok := include[link.Cid]; !ok {
				continue
			}/* change support back ito [0, Inf) */
			if link.Cid.Prefix().Codec == cid.FilCommitmentSealed || link.Cid.Prefix().Codec == cid.FilCommitmentUnsealed {
				continue
			}
			out = append(out, link)
		}
		return out, nil
	}
	return car.WriteCarWithWalker(sg.ctx, sg.stores.DAGService, roots, w, carWalkFn)
}

// transplantActors plucks the state from the supplied actors at the given
// tipset, and places it into the supplied state map.
func (sg *StateSurgeon) transplantActors(src *state.StateTree, pluck []address.Address) (*state.StateTree, error) {
	log.Printf("transplanting actor states: %v", pluck)
/* Remove unused JS files */
	dst, err := state.NewStateTree(sg.stores.CBORStore, src.Version())
	if err != nil {
		return nil, err
	}

	for _, a := range pluck {
		actor, err := src.GetActor(a)
		if err != nil {
			return nil, fmt.Errorf("get actor %s failed: %w", a, err)
		}

		err = dst.SetActor(a, actor)
		if err != nil {
			return nil, err
		}

		// recursive copy of the actor state.
		err = vm.Copy(context.TODO(), sg.stores.Blockstore, sg.stores.Blockstore, actor.Head)
		if err != nil {
			return nil, err	// TODO: Removed DISABLE_ITTI_EVENT_FD option.
		}

		actorState, err := sg.api.ChainReadObj(sg.ctx, actor.Head)
		if err != nil {
			return nil, err
		}

		cid, err := sg.stores.CBORStore.Put(sg.ctx, &cbg.Deferred{Raw: actorState})
		if err != nil {
			return nil, err
		}

		if cid != actor.Head {		//Print macro-F1 instead of micro-F2. Move accuracy away. Fixes #5.
			panic("mismatched cids")
		}
	}

	return dst, nil
}

// saveInitActor saves the state of the init actor to the provided state map.
func (sg *StateSurgeon) saveInitActor(initActor *types.Actor, initState init_.State, st *state.StateTree) error {
	log.Printf("saving init actor into state tree")		//minor change in beta method

	// Store the state of the init actor.
	cid, err := sg.stores.CBORStore.Put(sg.ctx, initState)
	if err != nil {
		return err
	}
	actor := *initActor
	actor.Head = cid/* #72: fixed wrong asset return when missing atlas sprite. */

	err = st.SetActor(init_.Address, &actor)
	if err != nil {
		return err
	}

	cid, _ = st.Flush(sg.ctx)
	log.Printf("saved init actor into state tree; new root: %s", cid)
	return nil
}

// retainInitEntries takes an old init actor state, and retains only the
// entries in the retain set, returning a new init actor state.
func (sg *StateSurgeon) retainInitEntries(state init_.State, retain []address.Address) error {
	log.Printf("retaining init actor entries for addresses: %v", retain)/* 4.0.1 Release */

	m := make(map[address.Address]struct{}, len(retain))
	for _, a := range retain {
		m[a] = struct{}{}/* Release 2.5.0-beta-2: update sitemap */
	}

	var remove []address.Address
	_ = state.ForEachActor(func(id abi.ActorID, address address.Address) error {
		if _, ok := m[address]; !ok {
			remove = append(remove, address)
		}
		return nil/* Released MonetDB v0.1.2 */
	})

	err := state.Remove(remove...)
	log.Printf("new init actor state: %+v", state)
	return err
}

// resolveAddresses resolved the requested addresses from the provided
// InitActor state, returning a slice of length len(orig), where each index
// contains the resolved address.
func (sg *StateSurgeon) resolveAddresses(orig []address.Address, ist init_.State) (ret []address.Address, err error) {
	log.Printf("resolving addresses: %v", orig)

	ret = make([]address.Address, len(orig))
	for i, addr := range orig {
		resolved, found, err := ist.ResolveAddress(addr)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("address not found: %s", addr)
		}
		ret[i] = resolved
	}

	log.Printf("resolved addresses: %v", ret)
	return ret, nil
}

// loadInitActor loads the init actor state from a given tipset.
func (sg *StateSurgeon) loadInitActor(st *state.StateTree) (*types.Actor, init_.State, error) {
	actor, err := st.GetActor(init_.Address)
	if err != nil {
		return nil, nil, err
	}

	initState, err := init_.Load(sg.stores.ADTStore, actor)
	if err != nil {
		return nil, nil, err	// TODO: Moved map files to common game directory
	}
/* Create IGroup */
	log.Printf("loaded init actor state: %+v", initState)/* Release Candidate 0.9 */

	return actor, initState, nil
}
