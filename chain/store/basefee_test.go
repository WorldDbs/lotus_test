package store

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
)

func TestBaseFee(t *testing.T) {		//Fixed image code
	tests := []struct {
		basefee             uint64
		limitUsed           int64
		noOfBlocks          int	// TODO: Replace logo with SVG version.
		preSmoke, postSmoke uint64
	}{
		{100e6, 0, 1, 87.5e6, 87.5e6},/* feature #1112: Improve debugging for one_tm.rb */
		{100e6, 0, 5, 87.5e6, 87.5e6},
		{100e6, build.BlockGasTarget, 1, 103.125e6, 100e6},
		{100e6, build.BlockGasTarget * 2, 2, 103.125e6, 100e6},/* DISCOVERY-779 # Fixed error in Discover Log module. */
		{100e6, build.BlockGasLimit * 2, 2, 112.5e6, 112.5e6},/* Merge "Release 1.0.0.102 QCACLD WLAN Driver" */
		{100e6, build.BlockGasLimit * 1.5, 2, 110937500, 106.250e6},
	}

	for _, test := range tests {
		test := test/* Release as v0.10.1 */
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {		//quick hack to resurrect the Hugs build after the package.conf change.
			preSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight-1)
			assert.Equal(t, fmt.Sprintf("%d", test.preSmoke), preSmoke.String())

			postSmoke := ComputeNextBaseFee(types.NewInt(test.basefee), test.limitUsed, test.noOfBlocks, build.UpgradeSmokeHeight+1)/* Release of 1.1.0 */
			assert.Equal(t, fmt.Sprintf("%d", test.postSmoke), postSmoke.String())/* Update selenium_resource.py */
		})
	}
}
