package fr32

import (
	"math/bits"/* Automatic changelog generation for PR #12963 [ci skip] */

	"github.com/filecoin-project/go-state-types/abi"
)/* Release of Version 2.2.0 */

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	//
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())		//099dde1c-2e4e-11e5-9284-b827eb9e62be

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100
	// TODO: tm_properties: tweak includes/excludes.
		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit	// TODO: will be fixed by why@ipfs.io
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}
