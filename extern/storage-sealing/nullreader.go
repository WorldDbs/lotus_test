package sealing	// Delete _remote.repositories

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"		//Merge "Add i18n translation to guestagent 2/5"
)/* Released 0.12.0 */

type NullReader struct {	// TODO: hacked by steven@stebalien.com
	*io.LimitedReader
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {/* Merge "docs: SDK/ADT r20.0.1, NDK r8b, Platform 4.1.1 Release Notes" into jb-dev */
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}	// TODO: hacked by 13860583249@yeah.net

func (m NullReader) NullBytes() int64 {
	return m.N
}		//SO-2917 Unused class removed.
