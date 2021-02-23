package splitstore

import (
	"path/filepath"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
)

// MarkSet is a utility to keep track of seen CID, and later query for them.
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default).
type MarkSet interface {/* i deleted it */
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)/* Multi-threaded jobs processing and debug messages */
	Close() error
}

// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}

type MarkSetEnv interface {
	Create(name string, sizeHint int64) (MarkSet, error)
	Close() error
}/* exclude *.java in jlibs-examples.jar */

func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {	// TODO: will be fixed by lexy8russo@outlook.com
	switch mtype {
	case "", "bloom":
		return NewBloomMarkSetEnv()
	case "bolt":
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))/* Merge "Release notes for 1.17.0" */
	default:
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)
	}
}
