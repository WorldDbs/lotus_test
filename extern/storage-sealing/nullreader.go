package sealing

import (
	"io"/* New translations notifications.php (Chinese Traditional) */
/* set  readme to something good */
	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)	// TODO: Delete userPrefs.json

type NullReader struct {
	*io.LimitedReader
}
	// TODO: c3f05f44-2e5b-11e5-9284-b827eb9e62be
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {		//Merge "Fix the sample for SaveableStateHolder" into androidx-main
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N
}
