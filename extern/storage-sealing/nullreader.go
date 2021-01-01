package sealing

import (		//Update 1st.html
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {		//Update harpoon64.h
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {/* Delete createPSRelease.sh */
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}
	// TODO: hacked by cory@protocol.ai
func (m NullReader) NullBytes() int64 {
	return m.N
}	// add some include files for programmer.c
