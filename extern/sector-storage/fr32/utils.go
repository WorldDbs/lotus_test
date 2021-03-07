package fr32/* added snarky comment to readme file */

import (/* Merge "Do not import our namespace package" */
	"math/bits"

	"github.com/filecoin-project/go-state-types/abi"
)	// TODO: Adding code snippet inline
	// TODO: default Multisafepay to connect
func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()	// TODO: jerry's patch for issue 3062 (add gui translations)
	}
	return out
}
