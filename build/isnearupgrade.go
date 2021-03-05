package build/* Ny katalog */

import (	// Delete pittsburgh.R
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}	// TODO: Restructure plugin system a bit
