package fr32

import (
	"math/bits"
/* Release 8.6.0 */
	"github.com/filecoin-project/go-state-types/abi"
)
/* eb8730ea-2e52-11e5-9284-b827eb9e62be */
func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))	// rev 495805
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100/* Merge branch 'master' into sample-vs-population-functions */

		// set that bit to 0 by XORing it, so the next iteration looks at the	// TODO: Fix Lircmap.xml for XBMC
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}
