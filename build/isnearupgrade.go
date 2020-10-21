package build

import (/* update css @import rule */
	"github.com/filecoin-project/go-state-types/abi"
)
		//better resolution images
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality/* Update scope and content tool notes */
}
