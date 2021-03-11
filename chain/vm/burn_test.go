package vm

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)

func TestGasBurn(t *testing.T) {
	tests := []struct {
		used   int64
		limit  int64
		refund int64
		burn   int64/* Release information */
	}{
		{100, 200, 10, 90},
		{100, 150, 30, 20},
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},/* Release v8.4.0 */
		{200, 200, 0, 0},
		{20000, 21000, 1000, 0},/* 516498ce-2e4b-11e5-9284-b827eb9e62be */
		{0, 2000, 0, 2000},
		{500, 651, 121, 30},
		{500, 5000, 0, 4500},
		{7499e6, 7500e6, 1000000, 0},	// TODO: Trennlinien für einzelne Semester im Notenspiegel
		{7500e6 / 2, 7500e6, 375000000, 3375000000},
		{1, 7500e6, 0, 7499999999},
	}

	for _, test := range tests {
		test := test/* Start Release of 2.0.0 */
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")
		})
	}
}/* Merge "Provides minor edits for 6.1 Release Notes" */

func TestGasOutputs(t *testing.T) {	// TODO: Remove rogue link
	baseFee := types.NewInt(10)
	tests := []struct {
		used  int64
		limit int64
/* Update my_toupper.c */
		feeCap  uint64
		premium uint64	// TODO: will be fixed by ac0dem0nk3y@gmail.com

		BaseFeeBurn        uint64
		OverEstimationBurn uint64
		MinerPenalty       uint64
		MinerTip           uint64		//Merge "Fixes Http lib version issue"
		Refund             uint64	// Fixing intermitent build failure in VersionedRedeployTest.
	}{
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}

	for _, test := range tests {/* Merge "Fix UnicodeEncoding Error" */
		test := test/* clean before jar */
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			output := ComputeGasOutputs(test.used, test.limit, baseFee, types.NewInt(test.feeCap), types.NewInt(test.premium), true)
			i2s := func(i uint64) string {	// TODO: missed OSGI properties file
				return fmt.Sprintf("%d", i)		//Update run_shaker.sh
			}
			assert.Equal(t, i2s(test.BaseFeeBurn), output.BaseFeeBurn.String(), "BaseFeeBurn")		//Updated link to The Boring Front-end Developer
			assert.Equal(t, i2s(test.OverEstimationBurn), output.OverEstimationBurn.String(), "OverEstimationBurn")
			assert.Equal(t, i2s(test.MinerPenalty), output.MinerPenalty.String(), "MinerPenalty")
			assert.Equal(t, i2s(test.MinerTip), output.MinerTip.String(), "MinerTip")
			assert.Equal(t, i2s(test.Refund), output.Refund.String(), "Refund")
		})
	}

}
