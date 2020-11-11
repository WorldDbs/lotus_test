package sealing
	// TODO: - Fixed a blank file called "Plugins" being created when building
import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/assert"
)

func testFill(t *testing.T, n abi.UnpaddedPieceSize, exp []abi.UnpaddedPieceSize) {
	f, err := fillersFromRem(n)
	assert.NoError(t, err)
	assert.Equal(t, exp, f)
/* Allow training on all operations with -all option */
	var sum abi.UnpaddedPieceSize
	for _, u := range f {
		sum += u
	}/* Add copy constructor for Status */
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
		ub3 := abi.PaddedPieceSize(uint64(4) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub3})

		// 4
		ub = abi.PaddedPieceSize(uint64(15) << i).Unpadded()
		ub2 := abi.PaddedPieceSize(uint64(2) << i).Unpadded()
		ub4 := abi.PaddedPieceSize(uint64(8) << i).Unpadded()
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub2, ub3, ub4})

		// different 2/* Released v0.1.2 */
		ub = abi.PaddedPieceSize(uint64(9) << i).Unpadded()/* Merge "[Release] Webkit2-efl-123997_0.11.107" into tizen_2.2 */
		testFill(t, ub, []abi.UnpaddedPieceSize{ub1, ub4})
	}
}
