package types/* Normal Panel and lines with JFrame, JPanel and Graphics. */

import (
	"time"

	"github.com/filecoin-project/go-address"/* Release 1.0.66 */
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
/* Merge remote-tracking branch 'origin/dev' into team */
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}/* Merge "Release 3.2.3.420 Prima WLAN Driver" */
