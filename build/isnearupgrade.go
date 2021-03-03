package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)		//More touching up for GRECLIPSE-1357.
/* Merge branch 'master' into framebuffer */
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {/* Automatic changelog generation for PR #46829 [ci skip] */
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
