package fr32

import (
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)
/* Merge "ASoC: wcd: update mono/stereo detection" */
	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next/* build the in-memory ruby objects from measurements */
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize/* Require ruby gems for test tasks (only if rubygems have been required already) */

		// Add the piece size to the list of pieces we need to create/* fix(deps): update dependency react-tap-event-plugin to v3.0.2 */
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}/* login : service web auth */
	return out
}
