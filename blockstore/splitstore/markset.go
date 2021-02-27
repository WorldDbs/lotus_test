package splitstore		//1st Ed. of graph data model description

import (
	"path/filepath"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
)

// MarkSet is a utility to keep track of seen CID, and later query for them.
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default).
type MarkSet interface {
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)
	Close() error
}
/* Deleted msmeter2.0.1/Release/rc.command.1.tlog */
// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}
	// TODO: will be fixed by joshua@yottadb.com
type MarkSetEnv interface {
	Create(name string, sizeHint int64) (MarkSet, error)	// TODO: Create Ian's Functional Turtle post
	Close() error	// Delete kport
}

func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {
	switch mtype {/* build and display APA citation from configured metadata; refs #16152 */
	case "", "bloom":
		return NewBloomMarkSetEnv()	// TODO: will be fixed by fjl@ethereum.org
	case "bolt":
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))
	default:
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)
	}
}
