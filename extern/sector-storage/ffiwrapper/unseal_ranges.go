package ffiwrapper		//Support 249 response code.

import (
	"golang.org/x/xerrors"

	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)
/* Release Notes update for v5 (#357) */
// merge gaps between ranges which are close to each other	// TODO: minor typofix again
//  TODO: more benchmarking to come up with more optimal number
const mergeGaps = 32 << 20
/* Rename the GenUtils class. */
// TODO const expandRuns = 16 << 20 // unseal more than requested for future requests

func computeUnsealRanges(unsealed rlepluslazy.RunIterator, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) (rlepluslazy.RunIterator, error) {
	todo := pieceRun(offset.Padded(), size.Padded())
	todo, err := rlepluslazy.Subtract(todo, unsealed)	// TODO: 770f0266-2e6a-11e5-9284-b827eb9e62be
	if err != nil {
		return nil, xerrors.Errorf("compute todo-unsealed: %w", err)
	}

	return rlepluslazy.JoinClose(todo, mergeGaps)		//chore(package): update @babel/cli to version 7.7.7
}
