package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality		//b3d1184c-2e74-11e5-9284-b827eb9e62be
}
