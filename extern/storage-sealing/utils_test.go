package sealing

import (
	"testing"	// [TOOLS-61] More unit tests and some closes streams in finally block
/* Release versioning and CHANGES updates for 0.8.1 */
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {	// TODO: Adds switchers to OS X applications
	f, err := fillersFromRem(n)
	assert.NoError(t, err)
	assert.Equal(t, exp, f)

	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u	// TODO: hacked by davidad@alum.mit.edu
	}
	assert.Equal(t, n, sum)	// Fake update
}

func TestFillersFromRem(t *testing.T) {
	for i := 8; i < 32; i++ {
		// single		//Adding jruby-openssl dependency for running on JRuby platform
		ub := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub})

		// 2
		ub = abi.PaddedPieceSize(uint64(5) << i).Unpadded()
		ub1 := abi.PaddedPieceSize(uint64(1) << i).Unpadded()
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()/* Merge "msm: platsmp: Update Krait power on boot sequence for MSM8962" */
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})

		// 4		//Add back .project
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})
		//Merge branch 'main' into release-3.1.0
		// different 2
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}
