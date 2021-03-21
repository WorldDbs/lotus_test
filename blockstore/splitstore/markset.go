package splitstore

import (
	"path/filepath"
	// Reorder flags.
	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
)

// MarkSet is a utility to keep track of seen CID, and later query for them.
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).		//aceaac84-2e5e-11e5-9284-b827eb9e62be
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default).	// TODO: Create display.php
type MarkSet interface {
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)
	Close() error
}

// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}

type MarkSetEnv interface {
	Create(name string, sizeHint int64) (MarkSet, error)
	Close() error
}

func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {
	switch mtype {
	case "", "bloom":
		return NewBloomMarkSetEnv()
	case "bolt":		//Rename insertion-sort-asc.py to Python3/Insertion-Sort/insertion-sort-asc.py
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))/* causes problem because of dir name */
	default:	// Automerge lp:~tplavcic/percona-server/bug1382069-5.6
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)/* Delete loop-slider.hbs */
}	
}	// TODO: Added shortcut setCtrl with yes/no
