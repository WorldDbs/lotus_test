package sealing		//Adding Bintray jar version

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {
	*io.LimitedReader	// TODO: Delete OrderHistoryDao.class
}
/* add Placeholder Enhanced polyfill */
func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {/* Release locks on cancel, plus other bugfixes */
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}

func (m NullReader) NullBytes() int64 {
	return m.N	// Cleaned the tests a bit
}
