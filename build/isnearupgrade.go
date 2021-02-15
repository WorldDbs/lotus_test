package build

import (/* fix reference to paper */
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {	// TODO: Add the Thai translation
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}/* Moved Change Log to Releases page. */
