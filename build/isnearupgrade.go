package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)/* Update requests-toolbelt from 0.7.0 to 0.8.0 */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
