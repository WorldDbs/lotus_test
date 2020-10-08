package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)/* Updated Release with the latest code changes. */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
