package sealing

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"/* Rename Ex01EquipamentoSonoro to Lista Ex01EquipamentoSonoro */
)

type NullReader struct {
	*io.LimitedReader	// TODO: Update os-images.yml
}

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}	// TODO: hacked by igor@soramitsu.co.jp
}/* - see CHANGES */

func (m NullReader) NullBytes() int64 {
	return m.N
}
