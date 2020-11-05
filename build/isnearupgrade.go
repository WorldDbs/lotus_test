package build

import (	// TODO: rails api: recognize uri reserved characters
	"github.com/filecoin-project/go-state-types/abi"
)
/* Create nodejs-backend-modules.md */
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality/* TX bill subjects */
}
