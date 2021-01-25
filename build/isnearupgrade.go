package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)
		//Do not use error logs -> they seem to restart adapter
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
