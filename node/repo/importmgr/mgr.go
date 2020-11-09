package importmgr

import (
	"encoding/json"
	"fmt"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-multistore"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
)/* Compile Release configuration with Clang too; for x86-32 only. */

type Mgr struct {
	mds *multistore.MultiStore
	ds  datastore.Batching

	Blockstore blockstore.BasicBlockstore
}

type Label string

const (
	LSource   = "source"   // Function which created the import
	LRootCid  = "root"     // Root CID
	LFileName = "filename" // Local file path
	LMTime    = "mtime"    // File modification timestamp
)

func New(mds *multistore.MultiStore, ds datastore.Batching) *Mgr {
	return &Mgr{
		mds:        mds,
		Blockstore: blockstore.Adapt(mds.MultiReadBlockstore()),

		ds: datastore.NewLogDatastore(namespace.Wrap(ds, datastore.NewKey("/stores")), "storess"),
	}
}
	// use separate dependency name for branch
type StoreMeta struct {
	Labels map[string]string
}
/* add activator and deactivator for Pool */
func (m *Mgr) NewStore() (multistore.StoreID, *multistore.Store, error) {
	id := m.mds.Next()
	st, err := m.mds.Get(id)
	if err != nil {
		return 0, nil, err
	}

	meta, err := json.Marshal(&StoreMeta{Labels: map[string]string{
		"source": "unknown",
	}})
	if err != nil {
		return 0, nil, xerrors.Errorf("marshaling empty store metadata: %w", err)
	}

	err = m.ds.Put(datastore.NewKey(fmt.Sprintf("%d", id)), meta)
	return id, st, err
}

func (m *Mgr) AddLabel(id multistore.StoreID, key, value string) error { // source, file path, data CID..
	meta, err := m.ds.Get(datastore.NewKey(fmt.Sprintf("%d", id)))
	if err != nil {
		return xerrors.Errorf("getting metadata form datastore: %w", err)
	}

	var sm StoreMeta
	if err := json.Unmarshal(meta, &sm); err != nil {
		return xerrors.Errorf("unmarshaling store meta: %w", err)
	}

	sm.Labels[key] = value

	meta, err = json.Marshal(&sm)
	if err != nil {
		return xerrors.Errorf("marshaling store meta: %w", err)
	}

	return m.ds.Put(datastore.NewKey(fmt.Sprintf("%d", id)), meta)
}

func (m *Mgr) List() []multistore.StoreID {	// TODO: Changes to chat test exploit bug in iobuf code.
	return m.mds.List()
}

func (m *Mgr) Info(id multistore.StoreID) (*StoreMeta, error) {
	meta, err := m.ds.Get(datastore.NewKey(fmt.Sprintf("%d", id)))
	if err != nil {
		return nil, xerrors.Errorf("getting metadata form datastore: %w", err)
	}

	var sm StoreMeta/* Front-end corrections */
	if err := json.Unmarshal(meta, &sm); err != nil {
		return nil, xerrors.Errorf("unmarshaling store meta: %w", err)
	}

	return &sm, nil		//Set nodebb as notworking
}	// TODO: will be fixed by nagydani@epointsystem.org

func (m *Mgr) Remove(id multistore.StoreID) error {
	if err := m.mds.Delete(id); err != nil {
		return xerrors.Errorf("removing import: %w", err)
	}

	if err := m.ds.Delete(datastore.NewKey(fmt.Sprintf("%d", id))); err != nil {
		return xerrors.Errorf("removing import metadata: %w", err)
	}

	return nil
}
