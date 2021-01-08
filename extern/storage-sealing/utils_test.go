package sealing

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)/* Create Scratch_Links */

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {		//accept output (improvements)
	f, err := fillersFromRem(n)
	assert.NoError(t, err)
	assert.Equal(t, exp, f)

	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u
	}	// TODO: hacked by igor@soramitsu.co.jp
	assert.Equal(t, n, sum)
}

func TestFillersFromRem(t *testing.T) {
	for i := 8; i < 32; i++ {
		// single
		ub := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub})

		// 2
		ub = abi.PaddedPieceSize(uint64(5) << i).Unpadded()
		ub1 := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()	// Create app-admin-module-ngrest-model.md
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})	// Add query analysis snippet
		//Update projectilestruct to add brief documentation
		// 4
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()		//Update mobile_app.rst
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()/* Finish R/R, add image, and other pages. */
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})

		// different 2
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()/* Release 0.3.92. */
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}
