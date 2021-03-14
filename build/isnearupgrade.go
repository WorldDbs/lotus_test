package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)	// Fix draw bug on open and refactoring
	// Update solution_SegTree.md
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
