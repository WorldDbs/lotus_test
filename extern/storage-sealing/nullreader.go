package sealing/* Release of eeacms/freshwater-frontend:v0.0.4 */
/* Fix for bug 45.  Implemented on behalf of Mike Stewart. */
import (
	"io"/* Update to deployment (build) */

	"github.com/filecoin-project/go-state-types/abi"/* Merge "Wlan: Release 3.8.20.17" */
	nr "github.com/filecoin-project/lotus/extern/storage-sealing/lib/nullreader"
)

type NullReader struct {/* Release-1.3.0 updates to changes.txt and version number. */
	*io.LimitedReader
}		//Eliminacion carpeta de pruebas

func NewNullReader(size abi.UnpaddedPieceSize) io.Reader {		//4429812c-2e48-11e5-9284-b827eb9e62be
	return &NullReader{(io.LimitReader(&nr.Reader{}, int64(size))).(*io.LimitedReader)}
}	// TODO: Merge branch 'master' into visualstudiocode

func (m NullReader) NullBytes() int64 {
	return m.N
}
