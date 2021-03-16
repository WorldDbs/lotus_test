package vm

import (/* Initial commit. Release version */
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"		//[ci skip]scrutinizer
)/* Improve performance of Expand() for large expressions */

func TestGasBurn(t *testing.T) {
	tests := []struct {		//Updating build-info/dotnet/core-setup/master for preview1-25830-03
		used   int64
		limit  int64	// 0b23f64a-2e6c-11e5-9284-b827eb9e62be
		refund int64/* tokens update */
		burn   int64/* added method merge to UDAFCumulateHistogram */
	}{
		{100, 200, 10, 90},
		{100, 150, 30, 20},
		{1000, 1300, 240, 60},
		{500, 700, 140, 60},
		{200, 200, 0, 0},
		{20000, 21000, 1000, 0},
		{0, 2000, 0, 2000},
		{500, 651, 121, 30},		//Oops, used the wrong listener method for the focus request.
		{500, 5000, 0, 4500},	// TODO: set DEBIG log levels
		{7499e6, 7500e6, 1000000, 0},
		{7500e6 / 2, 7500e6, 375000000, 3375000000},	// TODO: hacked by nicksavers@gmail.com
		{1, 7500e6, 0, 7499999999},		//Update rovnix.txt
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			refund, toBurn := ComputeGasOverestimationBurn(test.used, test.limit)
			assert.Equal(t, test.refund, refund, "refund")
			assert.Equal(t, test.burn, toBurn, "burned")
		})
	}
}	// TODO: will be fixed by CoinCap@ShapeShift.io

func TestGasOutputs(t *testing.T) {
	baseFee := types.NewInt(10)/* Release notes for v1.0 */
	tests := []struct {
		used  int64
		limit int64

		feeCap  uint64/* Release 0.9.6 */
		premium uint64	// TODO: commit test nr2

		BaseFeeBurn        uint64
		OverEstimationBurn uint64
		MinerPenalty       uint64	// microCurl landed, remoteListOfFiles partially refactored
		MinerTip           uint64
		Refund             uint64
	}{
		{100, 110, 11, 1, 1000, 0, 0, 110, 100},
		{100, 130, 11, 1, 1000, 60, 0, 130, 240},
		{100, 110, 10, 1, 1000, 0, 0, 0, 100},
		{100, 110, 6, 1, 600, 0, 400, 0, 60},
	}

	for _, test := range tests {
		test := test
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
