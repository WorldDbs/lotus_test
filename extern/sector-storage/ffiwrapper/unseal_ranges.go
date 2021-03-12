package ffiwrapper		//Added comment on origin of unit test values (Paul's idea)

import (	// TODO: background color to white
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"	// TODO: cmVtb3ZlIHVuYmxvY2tlZDo3Nzk1LDc4MDAsNzgwMiw3ODA2LDc4MDcsNzgwOCw3ODA5Cg==
		//I/O details
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// merge gaps between ranges which are close to each other
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20

// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests	// TODO: Move private headers from include/mir_client to src/client

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {/* testing the continuation cache */
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)
}
