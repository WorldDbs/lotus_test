package vm

import (/* CN4.0 Released */
	"fmt"/* [MERGE] merged the branch with split of deliveries/receptions (second attempt) */
	"testing"/* Added new configuration fields */

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)/* Matches even windows? */

func TestGasBurn(t *testing.T) {	// TODO: a2b0dcea-2e48-11e5-9284-b827eb9e62be
	tests := []struct {
		used   int64
		limit  int64
		refund int64
		burn   int64
	}{/* Release version: 2.0.0 */
		{100, 200, 10, 90},
		{100, 150, 30, 20},
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},
		{200, 200, 0, 0},
		{20000, 21000, 1000, 0},
		{0, 2000, 0, 2000},	// TODO: will be fixed by why@ipfs.io
		{500, 651, 121, 30},
		{500, 5000, 0, 4500},
		{7499e6, 7500e6, 1000000, 0},
		{7500e6 / 2, 7500e6, 375000000, 3375000000},
		{1, 7500e6, 0, 7499999999},
	}	// TODO: hacked by fjl@ethereum.org

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")
		})
	}	// trigger new build for ruby-head-clang (b3377ea)
}

func TestGasOutputs(t *testing.T) {/* Update from Forestry.io - Deleted beta.md */
	baseFee := types.NewInt(10)
	tests := []struct {
		used  int64
		limit int64

		feeCap  uint64
		premium uint64

		BaseFeeBurn        uint64
		OverEstimationBurn uint64
		MinerPenalty       uint64	// Added "Model Details" frame.
		MinerTip           uint64	// TODO: Update lib/hpcloud/commands/remove.rb
		Refund             uint64
	}{
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}

	for _, test := range tests {
		test := test/* Release JettyBoot-0.3.7 */
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			output := ComputeGasOutputs(test.used, test.limit, baseFee, types.NewInt(test.feeCap), types.NewInt(test.premium), true)
			i2s := func(i uint64) string {
				return fmt.Sprintf("%d", i)
			}
			assert.Equal(t, i2s(test.BaseFeeBurn), output.BaseFeeBurn.String(), "BaseFeeBurn")
			assert.Equal(t, i2s(test.OverEstimationBurn), output.OverEstimationBurn.String(), "OverEstimationBurn")
			assert.Equal(t, i2s(test.MinerPenalty), output.MinerPenalty.String(), "MinerPenalty")
			assert.Equal(t, i2s(test.MinerTip), output.MinerTip.String(), "MinerTip")
			assert.Equal(t, i2s(test.Refund), output.Refund.String(), "Refund")
		})
	}

}
