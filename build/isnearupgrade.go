package build	// TODO: TinyMCE: update to latest source including the WebKit backspace fix, see #23010

import (/* Add ErrorLog model to store errors */
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
